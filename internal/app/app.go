package app

import (
	"context"
	"github.com/ghn-rs/cloud-strife-user/internal/dependencies/config"
	"github.com/ghn-rs/corelib/src/logger"
	"log"
)

type App struct {
	Dependencies *Dependencies
	Repos        *Repos
	UseCases     *UseCases
}

func NewApp(ctx context.Context) (*App, error) {

	log.Println("Initializing Dependencies")
	dependencies := NewDependencies(ctx)

	if config.Config.Database.AutoMigrate {
		logger.Info(ctx, "Running database migrations")
		err := dependencies.DatabaseClient.RunMigrations()
		if err != nil {
			logger.Fatal(ctx, "Failed to run migrations: ", err)
		}
	}

	logger.Info(ctx, "Initializing Repositories")
	repos := NewRepos(dependencies)

	logger.Info(ctx, "Initializing UseCases")
	useCases := NewUseCases(repos)

	logger.Info(ctx, "App initialized")
	return &App{
		Dependencies: dependencies,
		Repos:        repos,
		UseCases:     useCases,
	}, nil
}
