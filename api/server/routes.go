package server

import (
	"net/http"

	"github.com/ablades/relevant/handlers"
	"github.com/labstack/echo/v4/middleware"
)

// SetRoutes will setup all the routes and
func (s *Server) SetRoutes() {
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch},
	}))

	handle := handlers.NewHandler(s.db)
	//client :=
	// Authentication Routes
	s.e.GET("/auth", handle.AuthReddit)
	s.e.GET("/authcallback", handle.AuthCallback)
	//s.e.GET("users/:name", handle.UserHome)

	//User
	// Routes
	s.e.POST("/users", handle.CreateUser)
	//s.e.GET("/users/:id", handle.getUser)
	//s.e.PUT("/users/:id", handle.updateUser)
	//s.e.DELETE("/users/:id", handle.deleteUser)
}
