package store

import "github.com/adarshhegde/backend-api-repo/internal/models"

// Store interface, all stores will implement this method for abstracting out
// the actual interaction with respective store
type Store interface {
	CreateUser(*models.User) error
	ListAllUsers() (error, []models.User)
}
