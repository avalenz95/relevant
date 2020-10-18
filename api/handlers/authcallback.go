package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

// AuthCallback for reddit authentication
func (h *Handler) AuthCallback(c echo.Context) (err error) {
	// Get Query Parameters
	code := c.QueryParam("code")

	// Custom http client for exchange call
	client := &http.Client{
		Transport: &oauth2.Transport{
			// configure token
			Source: h.config.TokenSource(oauth2.NoContext, &oauth2.Token{
				AccessToken: code,
			}),
			Base: h,
		},
	}

	ctx := context.WithValue(oauth2.NoContext, oauth2.HTTPClient, client)

	// Get token from auth code
	token, err := h.config.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	// Set client for token auto refresh
	h.client = h.config.Client(oauth2.NoContext, token)
	username := h.getRedditUserName()

	//return c.JSON(http.StatusOK, map[string]interface{}{"name": username})
	return c.Redirect(http.StatusTemporaryRedirect, "create/"+username)
}
