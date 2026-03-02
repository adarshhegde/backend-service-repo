package rpc

import (
	"context"

	backendservicepb "github.com/adarshhegde/backend-api-repo/proto-files/generated-code/backendservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (handler *gRPCHandler) ListAllUsers(context.Context, *emptypb.Empty) (*backendservicepb.ListAllUsersResponse, error) {
	err, users := handler.InternalServices.GetUserSvc().ListAllUsers()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal Error")
	}

	usersProto := users.ToProto() // hehe magic!

	return &backendservicepb.ListAllUsersResponse{Users: usersProto}, nil
}
