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

	route := handlers.NewHandler(s.db)

	s.e.GET("/auth", route.RedditAuth)
	s.e.GET("/authcallback", route.AuthCallback)
}
