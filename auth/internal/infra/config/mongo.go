// Package config
package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(env *Env) (*mongo.Client, error) {
	mongoRL := env.MongoDBConnectionString

	if mongoRL == "" {
		panic("MongoDB connection string is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoRL))
	if err != nil {
		return nil, err
	}

	return client, nil
}
