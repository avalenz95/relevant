package handlers

import (
	"context"
	"fmt"
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
			Source: h.config.TokenSource(c.Request().Context(), &oauth2.Token{
				AccessToken: code,
			}),
			Base: h,
		},
	}

	ctx := context.WithValue(c.Request().Context(), oauth2.HTTPClient, client)

	// Get token from auth code
	token, err := h.config.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	// Set client for token auto refresh
	h.client = h.config.Client(c.Request().Context(), token)
	fmt.Println("Client is set! Reddit API Requests can be made!")

	userinfo := h.request(c, http.MethodGet, "https://oauth.reddit.com/api/v1/me", nil)
	// Create user

	return c.JSONBlob(http.StatusOK, userinfo)
}
