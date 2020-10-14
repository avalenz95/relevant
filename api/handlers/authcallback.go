package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ablades/relevant/config"
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
	// set context
	ctx := c.Request().Context()
	// csrf check
	if queryState != state {
		return c.String(http.StatusForbidden, "csrf detected")
	}

	authConfig := config.GetAuthConfig()
	token, err := authConfig.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	// Custom http client
	*h.client = http.Client{
		Transport: &oauth2.Transport{
			// configure token
			Source: authConfig.TokenSource(ctx, token),
			Base:   h,
		},
	}

	resp, err := h.client.Get("https://oauth.reddit.com/api/v1/me.json")
	fmt.Println(resp.Header.Get("User-Agent"))
	content, err := ioutil.ReadAll(resp.Body)

	return c.JSONBlob(http.StatusOK, content)
}
