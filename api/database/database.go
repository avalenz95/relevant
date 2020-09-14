package database

import (
	"github.com/ablades/relevant/api/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func Connect() *mongo.Database {
	dbConfig := config.GetDBConfig()
}
