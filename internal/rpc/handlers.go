package rpc

import (
	"github.com/adarshhegde/backend-api-repo/internal/services"
	backendservicepb "github.com/adarshhegde/backend-api-repo/proto-files/generated-code/backendservice"
)

type gRPCHandler struct {
	backendservicepb.UnimplementedBackendServiceServer
	InternalServices services.InternalServices
}

func New(opts services.InternalServicesOpts) backendservicepb.BackendServiceServer {
	return &gRPCHandler{
		InternalServices: services.New(opts),
	}
}
