package main

import (
	"context"
	"github.com/ghn-rs/cloud-strife-user/internal/app"
	"github.com/ghn-rs/corelib/src/logger"
	"log"
)

func main() {

	log.Println("Starting cloud-strife-user")
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	log.Println("Initializing App")
	application, err := app.NewApp(ctx)
	if err != nil {
		panic(err)
	}

	logger.Info(ctx, "Creating servers")
	servers, err := application.NewServers()
	if err != nil {
		panic(err)
	}

	logger.Info(ctx, "Starting servers")
	if err := servers.Start(ctx); err != nil {
		panic(err)
	}

	<-ctx.Done()
}
