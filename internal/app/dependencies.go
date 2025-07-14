package app

import (
	"context"
	"github.com/ghn-rs/cloud-strife-user/internal/dependencies"
	"github.com/ghn-rs/cloud-strife-user/internal/dependencies/config"
	"github.com/ghn-rs/corelib/src/logger"
	"github.com/sirupsen/logrus"
	"log"
)

type Dependencies struct {
	DatabaseClient *dependencies.DatabaseClient
}

func NewDependencies(ctx context.Context) *Dependencies {

	log.Println("Loading Config")
	config.InitConfig()
	logLevel, err := logrus.ParseLevel(config.Config.Elastic.Level)
	if err != nil {
		logLevel = logrus.DebugLevel
	}

	log.Println("Initializing Logger")
	logger.Init(logger.LoggerConfig{
		Level:            logLevel,
		ElasticsearchURL: []string{config.Config.Elastic.Url},
		ServiceName:      "cloud-strife-user",
	})

	logger.Info(ctx, "Connecting to Database")
	databaseClient, err := dependencies.NewDatabaseClient()
	if err != nil {
		logger.Fatal(ctx, "Failed to connect to database: ", err)
	}

	logger.Info(ctx, "Dependencies initialized")
	return &Dependencies{
		DatabaseClient: databaseClient,
	}
}
