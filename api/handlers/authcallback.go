package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"
)

// AuthCallback for reddit authentication
func (h *Handler) AuthCallback(c echo.Context) (err error) {
	// Get Query Parameters
	code := c.QueryParam("code")
	queryState := c.QueryParam("state")
	state := c.Get(middleware.DefaultCSRFConfig.ContextKey)

	// csrf check
	if queryState != state {
		return c.String(http.StatusForbidden, "csrf detected")
	}

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
	h.client = h.config.Client(context.Background(), token)
	fmt.Println("Client is set! Reddit API Requests can be made!")

	content := h.request(http.MethodGet, "https://oauth.reddit.com/api/v1/me")

	return c.String(http.StatusOK, string(content))
}
