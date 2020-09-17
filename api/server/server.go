package server

import (
	"github.com/labstack/echo/v4"
	"github.com/tempor1s/echo-boiler/db"
	"go.mongodb.org/mongo-driver/mongo"
)

// Server represents components of a server
type Server struct {
	e  *echo.Echo
	db *mongo.Database
}

// NewServer creates a Server and instantiates the DB if not provided
func NewServer(database *mongo.Database) *Server {
	if database == nil {
		database = db.Connect()
	}

	return &Server{
		e:  echo.New(),
		db: database,
	}
}

// Start the server
func (s *Server) Start(port string) {
	// register routes
	s.SetRoutes()
	s.e.Logger.Fatal(s.e.Start(port))
}

// Close stops the Server
func (s *Server) Close() {
	// stop the server
	s.e.Close()
}
