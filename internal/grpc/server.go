package grpc

import (
	"context"
	"github.com/ghn-rs/cloud-strife-user/internal/usecase/userusecase"
	cloudstrife "github.com/ghn-rs/cloud-strife-user/proto/gen"
	"github.com/ghn-rs/corelib/src/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Server struct {
	cloudstrife.UnimplementedUserServiceServer

	UserUseCase userusecase.UseCase
}

func (s *Server) Start(ctx context.Context, address string) error {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	cloudstrife.RegisterUserServiceServer(server, s)
	reflection.Register(server)

	logger.Info(ctx, "Starting gRPC server on ", address)
	return server.Serve(lis)
}
