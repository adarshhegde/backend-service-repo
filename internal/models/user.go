package models

import (
	backendservicepb "github.com/adarshhegde/backend-api-repo/proto-files/generated-code/backendservice"
)

// Base Model, intended to be the default struct used in http api
type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// ToProto is a receiver function that acts as a helper for converting the regular model
// to it's gRPC Equivalant protobuf generated struct.
// Instead of manually converting it, this helper function makes life easier.
// The methods in the Store will always return a models Base struct, and the same can be
// used in gRPC, by simply calling ToProto() on the object, which gives the pb struct in return.
func (user User) ToProto() *backendservicepb.User {
	return &backendservicepb.User{
		Username: user.Username,
		Password: user.Password,
	}
}

// List of users
type Users []User

// We can have receiver functions on any custom type, even an Array/slice!
func (users Users) ToProto() []*backendservicepb.User {
	protoUsers := make([]*backendservicepb.User, len(users))
	for i, u := range users {
		protoUsers[i] = u.ToProto()
	}

	return protoUsers
}
