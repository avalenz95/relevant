package db

import (
	"context"
	"log"

	"github.com/ablades/relevant/api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Database {
	dbConfig := config.GetDBConfig()

	auth := options.Credential{
		Username: dbConfig.Username,
		Password: dbConfig.Password,
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbConfig.URI).SetAuth(auth))
	if err != nil {
		log.Fatal(err)
	}

	// TODO: change this
	return client.Database("boiler")
}
