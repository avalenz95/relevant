package db

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connect to a given database
func Connect() *mongo.Database {
	// Build URI
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
		viper.Get("db.user"),
		viper.Get("db.password"),
		viper.Get("db.cluster"),
		viper.Get("db.name"),
	)

	// Close on failed connection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	// TODO: Send to logs
	fmt.Println("Successfully connected and pinged.")
	fmt.Println(client.ListDatabaseNames(context.TODO(), bson.D{}))
	fmt.Println(client.Database(viper.GetString("db.name")).ListCollectionNames(context.TODO(), bson.D{}))

	return client.Database(viper.GetString("db.name"))
}
