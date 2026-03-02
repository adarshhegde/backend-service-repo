package models

import (
	backendservicepb "github.com/adarshhegde/backend-api-repo/proto-files/generated-code/backendservice"
)

// Base Model, can be used in HTTP API By Default
type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// Adding a receiver function to the model type
// this ToProto helper allows you to easily convert a
// model object instance to a Protobuf instance by simply calling ToProto() on the object
func (user User) ToProto() *backendservicepb.User {
	return &backendservicepb.User{
		Username: user.Username,
		Password: user.Password,
	}
}

// List of users
type Users []User

// Same type of helper as above.
func (users Users) ToProto() []*backendservicepb.User {
	protoUsers := make([]*backendservicepb.User, len(users))
	for i, u := range users {
		protoUsers[i] = u.ToProto()
	}

	return protoUsers
}
