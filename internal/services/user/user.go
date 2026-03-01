package user

import (
	"github.com/adarshhegde/backend-api-repo/internal/cmd/http/dto"
	"github.com/adarshhegde/backend-api-repo/internal/models"
	"github.com/adarshhegde/backend-api-repo/internal/store"
)

type UserService interface {
	CreateUser(dto.CreateUserRequest) error
	ListAllUsers() (error, []models.User)
}

type UserServiceImpl struct {
	Store store.Store
}

func (svc *UserServiceImpl) CreateUser(request dto.CreateUserRequest) error {
	return svc.Store.CreateUser(&models.User{
		Username: request.Username,
		Password: request.Password,
	})
}
func (svc *UserServiceImpl) ListAllUsers() (error, []models.User) {
	return svc.Store.ListAllUsers()
}
