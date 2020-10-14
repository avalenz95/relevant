package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// AuthCallback handler for reddit authentication
func (h *Handler) AuthCallback(c echo.Context) (err error) {
	// Get Query Parameters
	queryState := c.QueryParam("state")
	//code := c.QueryParam("code")
	state := c.Get(middleware.DefaultCSRFConfig.ContextKey)

	if queryState != state {
		return c.String(http.StatusForbidden, "csrf detected")
	}

	// TODO: add token fetch
	return c.String(http.StatusOK, "validated :)")
}
