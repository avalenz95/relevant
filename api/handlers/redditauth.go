package handlers

import (
	"fmt"
	"net/http"

	"github.com/ablades/relevant/config"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

// RedditAuth route
func (h *Handler) RedditAuth(c echo.Context) (err error) {

	authConfig := config.GetAuthConfig()
	fmt.Print(authConfig)

	// Set state and additional params
	url := authConfig.AuthCodeURL(uuid.New().String(),
		oauth2.SetAuthURLParam("response_type", "code"),
		oauth2.SetAuthURLParam("duration", "temporary"),
	)

	return c.Redirect(http.StatusTemporaryRedirect, url)
}
