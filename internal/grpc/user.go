package grpc

import (
	"context"
	pb "github.com/ghn-rs/cloud-strife-user/proto/gen"
	coreproto "github.com/ghn-rs/corelib/proto/gen"
	grpc "github.com/ghn-rs/corelib/src/grpc"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (resp *pb.UserResponse, err error) {
	return grpc.HandleGrpc(ctx, req, resp, s.UserUseCase.CreateUser)
}

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (resp *pb.UserResponse, err error) {
	return grpc.HandleGrpc(ctx, req, resp, s.UserUseCase.GetUser)
}

func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (resp *pb.UserResponse, err error) {
	return grpc.HandleGrpc(ctx, req, resp, s.UserUseCase.UpdateUser)
}

func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (resp *coreproto.GenericResponse, err error) {
	return grpc.HandleGrpc(ctx, req, resp, s.UserUseCase.DeleteUser)
}

func (s *Server) PaginateUser(ctx context.Context, req *pb.PaginateUserRequest) (resp *pb.UsersResponse, err error) {
	return grpc.HandleGrpc(ctx, req, resp, s.UserUseCase.PaginateUser)
}

func (s *Server) ListUsers(ctx context.Context, req *pb.EmptyRequest) (resp *pb.UsersResponse, err error) {
	return grpc.HandleGrpc(ctx, req, resp, s.UserUseCase.ListUsers)
}

func (s *Server) HealthCheck(ctx context.Context, req *pb.EmptyRequest) (resp *coreproto.GenericResponse, err error) {
	return grpc.HandleGrpc(ctx, req, resp, s.UserUseCase.HealthCheck)
}
