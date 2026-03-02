package main

import (
	"fmt"
	"log"
	"net"

	"github.com/adarshhegde/backend-api-repo/internal/config"
	"github.com/adarshhegde/backend-api-repo/internal/rpc"
	"github.com/adarshhegde/backend-api-repo/internal/services"
	backendservicepb "github.com/adarshhegde/backend-api-repo/proto-files/generated-code/backendservice"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	config, err := config.Load()
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// connect to DBs
	client, err := mongo.Connect(options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		panic(err)
	}

	//common initialisation for all dependencies
	opts := services.InternalServicesOpts{
		MongoClient: client,
	}

	rpcHandler := rpc.New(opts)

	backendservicepb.RegisterBackendServiceServer(s, rpcHandler)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
