package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect to a given database
func Connect() *mongo.Database {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	fmt.Println("here")
	fmt.Println(viper.GetString("reddit.username"))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(viper.GetString("db.testuri")))
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
