package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/adarshhegde/backend-api-repo/internal/api"
	"github.com/adarshhegde/backend-api-repo/internal/config"
	"github.com/adarshhegde/backend-api-repo/internal/services"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	config, err := config.Load()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.RequestLogger())

	// connect to DBs
	client, err := mongo.Connect(options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		panic(err)
	}

	//common initialisation for all dependencies
	opts := services.InternalServicesOpts{
		MongoClient: client,
	}

	handler := api.New(opts)

	// Connect Routes with their Handlers :D
	// all of them will be able to access these internal services defined above
	// using dependency injection done via the interface technique.
	e.POST("/user/create", handler.CreateUser)
	e.GET("/user/all", handler.ListAllUsers)

	// listening for the interrupt / kill signal to enable graceful shutdown of the server
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	sc := echo.StartConfig{
		Address:         fmt.Sprintf(":%d", config.HTTPPort),
		GracefulTimeout: 5 * time.Second, // max time to allow for requests to flush out, and then kill server
	}
	if err := sc.Start(ctx, e); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}

}
