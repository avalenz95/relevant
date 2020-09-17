package handlers

import "go.mongodb.org/mongo-driver/mongo"

// Handler is used so that we can dependency inject different services into the database
type Handler struct {
	db *mongo.Database
}

// New will create a new handler with the given databse and other dependencies that your routes may need
func NewHandler(db *mongo.Database) *Handler {
	return &Handler{
		db,
	}
}
