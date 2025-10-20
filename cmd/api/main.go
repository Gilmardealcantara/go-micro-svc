package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Gilmardealcantara/go-micro-svc/pkg/config"
	"github.com/Gilmardealcantara/go-micro-svc/pkg/tel"
	"github.com/Gilmardealcantara/go-micro-svc/pkg/tel/log"
	"github.com/joho/godotenv"

	"github.com/Gilmardealcantara/go-micro-svc/db"
	docs "github.com/Gilmardealcantara/go-micro-svc/docs/swagger"
	"github.com/Gilmardealcantara/go-micro-svc/pkg/api"
	"github.com/Gilmardealcantara/go-micro-svc/pkg/server"
)

//	@title			Hello World App
//	@version		1.0
//	@description	This is the API for the sophisticated Hello World App

//	@BasePath /

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	cfg := config.NewConfig()
	ctx := context.Background()

	// initializer newrelic instrumentation and structured log
	slogInstance, err := tel.InitializeWithSlog(cfg)
	if err != nil {
		panic(err)
	}
	logger := log.New(slogInstance)

	// Set host dynamically for swagger docs
	docs.SwaggerInfo.Host = cfg.SwaggerAppURL

	// once a DB is provisioned you can connect to it like so
	connectedDB, err := db.Conn(ctx, cfg)
	if err != nil {
		// TODO: only for success template deploy, replace to fatal,
		logger.Error(ctx, err.Error())
	}

	if cfg.BindPort == "" {
		cfg.BindPort = "8000"
	}

	port := fmt.Sprintf(":%s", cfg.BindPort)
	router := api.SetupRouter(&cfg, connectedDB, logger)

	newSvr := &http.Server{
		Addr:    port,
		Handler: router,
	}

	// You can pass additional shutdown functions to the server to be executed
	// upon graceful shutdown.
	svrWithShutdown := server.New(
		logger,
		newSvr,
		server.WithAdditionalShutdown(db.ClosePool(connectedDB)),
	)

	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	defer stop()

	svrWithShutdown.Start(ctx)
}
