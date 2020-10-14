package handlers

import (
	"fmt"
	"net/http"

	"github.com/ablades/relevant/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"
)

// RedditAuth route
func (h *Handler) RedditAuth(c echo.Context) (err error) {
	authConfig := config.GetAuthConfig()
	fmt.Print(authConfig)
	c.Echo().Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLength: 25,
		TokenLookup: "query:csrf",
	}))

	state := c.Get(middleware.DefaultCSRFConfig.ContextKey).(string)
	// Set state and additional params
	url := authConfig.AuthCodeURL(state,
		oauth2.SetAuthURLParam("response_type", "code"),
		oauth2.SetAuthURLParam("duration", "temporary"),
	)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}
