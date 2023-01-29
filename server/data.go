package server

import (
	"context"

	schemapb "github.com/iptecharch/schema-server/protos/schema_server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// datastore
// func (s *Server) GetDataStore(ctx context.Context, req *schemapb.GetDataStoreRequest) (*schemapb.GetDataStoreResponse, error) {
// 	name := req.GetName()
// 	if name == "" {
// 		return nil, status.Error(codes.InvalidArgument, "missing name attribute")
// 	}
// 	s.md.RLock()
// 	defer s.md.RUnlock()
// 	ds, ok := s.datastores[name]
// 	if !ok {
// 		return nil, status.Errorf(codes.InvalidArgument, "unknown datastore %s", name)
// 	}
// 	cands := ds.Candidates()
// 	rsp := &schemapb.GetDataStoreResponse{
// 		Datastore: make([]*schemapb.DataStore, 0, len(cands)),
// 	}

// 	for _, cand := range cands {
// 		rsp.Datastore = append(rsp.Datastore,
// 			&schemapb.DataStore{
// 				Type: *schemapb.Type_CANDIDATE.Enum(),
// 				Name: cand,
// 			},
// 		)
// 	}

// 	return rsp, nil
// }

// func (s *Server) CreateDataStore(ctx context.Context, req *schemapb.CreateDataStoreRequest) (*schemapb.CreateDataStoreResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method CreateDataStore not implemented")
// }

// func (s *Server) DeleteDataStore(ctx context.Context, req *schemapb.DeleteDataStoreRequest) (*schemapb.DeleteDataStoreResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method DeleteDataStore not implemented")
// }

// func (s *Server) Commit(ctx context.Context, req *schemapb.CommitRequest) (*schemapb.CommitResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
// }

// func (s *Server) Discard(ctx context.Context, req *schemapb.DiscardRequest) (*schemapb.DiscardResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method Discard not implemented")
// }

// data

func (s *Server) GetData(ctx context.Context, req *schemapb.GetDataRequest) (*schemapb.GetDataResponse, error) {
	logrus.Debugf("received GetDataRequest: %v", req)
	name := req.GetName()
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing datastore name")
	}
	s.md.RLock()
	defer s.md.RUnlock()
	ds, ok := s.datastores[name]
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "unknown datastore %s", name)
	}
	return ds.Get(ctx, req)
}

func (s *Server) SetData(ctx context.Context, req *schemapb.SetDataRequest) (*schemapb.SetDataResponse, error) {
	logrus.Debugf("received SetDataRequest: %v", req)
	name := req.GetName()
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing datastore name")
	}
	s.md.RLock()
	defer s.md.RUnlock()
	ds, ok := s.datastores[name]
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "unknown datastore %s", name)
	}
	return ds.Set(ctx, req)
}

func (s *Server) Diff(ctx context.Context, req *schemapb.DiffRequest) (*schemapb.DiffResponse, error) {
	logrus.Debugf("received DiffRequest: %v", req)
	name := req.GetName()
	if name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "missing datastore name")
	}
	return nil, status.Errorf(codes.Unimplemented, "method Diff not implemented")
}

func (s *Server) Subscribe(req *schemapb.SubscribeRequest, stream schemapb.DataServer_SubscribeServer) error {
	logrus.Debugf("received SubscribeRequest: %v", req)
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
