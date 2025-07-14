package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/ghn-rs/cloud-strife-user/internal/dependencies/config"
	"github.com/ghn-rs/cloud-strife-user/internal/grpc"
	"github.com/ghn-rs/corelib/src/logger"
	"golang.org/x/sync/errgroup"
	"strconv"
)

type Servers struct {
	grpcServer *grpc.Server
}

func (app *App) NewServers() (*Servers, error) {

	useCases := app.UseCases

	grpcServer := &grpc.Server{
		UserUseCase: useCases.User,
	}

	return &Servers{
		grpcServer: grpcServer,
	}, nil
}

func (s *Servers) Start(ctx context.Context) error {
	var g errgroup.Group

	g.Go(func() error {
		logger.Info(ctx, "Starting gRPC server")
		address := config.Config.App.Address + ":" + strconv.Itoa(config.Config.App.Port)
		if err := s.grpcServer.Start(ctx, address); err != nil {
			return fmt.Errorf("failed to start gRPC server: %w", err)
		}
		return nil
	})

	return g.Wait()
}

func (s *Servers) Stop() error {
	return errors.New("stop not implemented")
}
