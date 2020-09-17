package handlers

import "go.mongodb.org/mongo-driver/mongo"

// Handler is used so that we can dependency inject different services into the database
type Handler struct {
	db *mongo.Database
}

// NewHandler with given databse and other dependencies
func NewHandler(db *mongo.Database) *Handler {
	return &Handler{
		db,
	}
}
