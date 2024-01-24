package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"

	"github.com/iptecharch/schema-server/pkg/config"
	"github.com/iptecharch/schema-server/pkg/server"
)

var version = "dev"
var commit = ""

var configFile string
var debug bool
var trace bool
var stop bool
var versionFlag bool

func main() {
	pflag.StringVarP(&configFile, "config", "c", "schema-server.yaml", "config file path")
	pflag.BoolVarP(&debug, "debug", "d", false, "set log level to DEBUG")
	pflag.BoolVarP(&trace, "trace", "t", false, "set log level to TRACE")
	pflag.BoolVarP(&versionFlag, "version", "v", false, "print version")
	pflag.Parse()

	if versionFlag {
		fmt.Printf("%s-%s\n", version, commit)
		return
	}

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetLevel(log.InfoLevel)
	if debug {
		log.SetLevel(log.DebugLevel)
	}
	if trace {
		log.SetLevel(log.TraceLevel)
	}
	var s *server.Server
START:
	if s != nil {
		s.Stop()
	}
	cfg, err := config.New(configFile)
	if err != nil {
		log.Errorf("failed to read config: %v", err)
		os.Exit(1)
	}
	b, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		log.Errorf("failed to marshal config: %v", err)
		os.Exit(1)
	}
	log.Infof("read config:\n%s", string(b))

	s, err = server.NewServer(cfg)
	if err != nil {
		log.Errorf("failed to create server: %v", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	setupCloseHandler(cancel)

	err = s.Serve(ctx)
	if err != nil {
		if stop {
			return
		}
		log.Errorf("failed to run server: %v", err)
		time.Sleep(time.Second)
		goto START
	}
}

func setupCloseHandler(cancelFn context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-c
		fmt.Fprintf(os.Stderr, "\nreceived signal '%s'. terminating...\n", sig.String())
		stop = true
		cancelFn()
		time.Sleep(500 * time.Millisecond)
		os.Exit(0)
	}()
}
