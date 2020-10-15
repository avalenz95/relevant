package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"
)

// AuthReddit route
func (h *Handler) AuthReddit(c echo.Context) (err error) {
	// create csrf token
	c.Echo().Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLength: 25,
		TokenLookup: "query:csrf",
	}))

	// Set state and additional params
	state := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
	url := h.config.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("response_type", "code"),
		oauth2.SetAuthURLParam("duration", "temporary"),
	)

	return c.Redirect(http.StatusTemporaryRedirect, url)
}
