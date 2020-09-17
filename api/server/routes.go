package server

import (
	"github.com/ablades/relevant/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Routes will setup all the routes and
func (s *Server) SetRoutes() {
	s.e.Use(middleware.Logger())
	// s.e.Use(middleware.Recover())
	s.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE, echo.PATCH},
	}))

	r := handlers.NewHandler(s.db)

	s.e.GET("/", r.Hello)
	// Setup all your routes here!
}
