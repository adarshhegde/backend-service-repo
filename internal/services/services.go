package services

import (
	"github.com/adarshhegde/backend-api-repo/internal/services/user"
	"github.com/adarshhegde/backend-api-repo/internal/store/mongodb"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type InternalServices interface {
	GetUserSvc() user.UserService
}

type InternalServicesImpl struct {
	userSvc user.UserService
}

type InternalServicesOpts struct {
	MongoClient *mongo.Client
}

func New(opts InternalServicesOpts) InternalServices {
	return InternalServicesImpl{
		userSvc: &user.UserServiceImpl{
			Store: mongodb.New(opts.MongoClient),
		},
	}
}

func (svcs InternalServicesImpl) GetUserSvc() user.UserService {
	return svcs.userSvc
}
