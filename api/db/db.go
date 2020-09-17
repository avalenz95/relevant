package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ablades/relevant/api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect to a given database
func Connect() *mongo.Database {
	dbConfig := config.GetDBConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbConfig.URI))
	if err != nil {
		log.Fatal(err)
	}

	// Print Active Collections
	collections, _ := client.Database("testdb").ListCollectionNames(ctx, nil)
	for _, coll := range collections {
		fmt.Printf("Connected to Collection: %s", coll)
	}

	return client.Database("testdb")
}

// ActiveCollections current collections a
func ActiveCollections(db *mongo.Database) {

	// use a filter to only select capped collections
	result, err := db.ListCollectionNames(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, coll := range result {
		fmt.Println(coll)
	}
}
