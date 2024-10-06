package auth

import (
	"context"
	cadiav1 "github.com/HorusEternal/gate_protos/gen/go/cadia"
	"google.golang.org/grpc"
)

type serverApi struct {
	cadiav1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	cadiav1.RegisterAuthServer(gRPC, &serverApi{})
}

func (s *serverApi) Login(ctx context.Context, req *cadiav1.LoginRequest) (*cadiav1.LoginResponse, error) {
	panic("Implement me")
}

func (s *serverApi) Register(ctx context.Context, req *cadiav1.RegisterRequest) (*cadiav1.RegisterResponse, error) {
	panic("Implement me")
}

func (s *serverApi) IsAdmin(ctx context.Context, req *cadiav1.IsAdminRequest) (*cadiav1.IsAdminResponse, error) {
	panic("Implement me")
}
