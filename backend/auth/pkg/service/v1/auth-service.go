package v1

import (
	"context"

	v1 "auth/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVersion = "v1"
)

type GRPCServer struct{}

// create auth service for testing
func NewAuthServiceServer() v1.AuthServiceServer {
	return &GRPCServer{}
}

func (s *GRPCServer) checkAPI(api string) error {
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

func (s *GRPCServer) ReadAll(ctx context.Context, req *v1.ReadAllRequest) (*v1.
	ReadAllResponse, error) {

	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	return &v1.ReadAllResponse{
		Api:  "v1",
		Auth: nil,
	}, nil
}

func (s *GRPCServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error) {

	// check if the APi version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	return &v1.ReadResponse{
		Api:  "v1",
		Auth: nil,
	}, nil
}

func (s *GRPCServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {

	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	return &v1.CreateResponse{
		Api: "v1",
		Id:  "hello world",
	}, nil
}

func (s *GRPCServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {

	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	return &v1.DeleteResponse{
		Api:     "v1",
		Deleted: 1,
	}, nil
}

func (s *GRPCServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {

	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	return &v1.UpdateResponse{
		Api:     "v1",
		Updated: 1,
	}, nil
}
