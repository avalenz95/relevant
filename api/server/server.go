package server

import (
	"context"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

// Server represents components of a server
type Server struct {
	e  *echo.Echo
	db *mongo.Database
}

// NewServer create create db and instantiate server
func NewServer(db *mongo.Database) *Server {

	if db == nil {
		db = mongo.Connect(context.TODO())
	}

	server := Server{
		e:  echo.New(),
		db: db,
	}

	return server
}
