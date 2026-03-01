package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/adarshhegde/backend-api-repo/internal/models"
	"github.com/adarshhegde/backend-api-repo/internal/store"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoStore struct {
	Client *mongo.Client
}

func New(client *mongo.Client) store.Store {
	return &MongoStore{
		Client: client,
	}
}

// Implementations of the Store
func (ms *MongoStore) CreateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := ms.Client.Database("db").Collection("user")
	result, err := coll.InsertOne(ctx, user)
	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

	return err
}

func (ms *MongoStore) ListAllUsers() (error, []models.User) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := ms.Client.Database("db").Collection("user")

	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return err, nil
	}

	var results []models.User
	if err = cursor.All(ctx, &results); err != nil {
		return err, nil
	}

	return nil, results
}
