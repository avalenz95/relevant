package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ablades/relevant/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
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

	authConfig := config.GetAuthConfig()
	// token, err := authConfig.Exchange(ctx, code)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Custom http client
	client := &http.Client{
		Transport: &oauth2.Transport{
			// configure token
			Source: authConfig.TokenSource(oauth2.NoContext, &oauth2.Token{
				AccessToken: code,
			}),
			Base: &transp{
				config:    authConfig,
				userAgent: fmt.Sprintf("relevant_for_reddit/0.0 (by /u/%s)", viper.GetString("reddit.username")),
			},
		},
	}

	ctx := context.WithValue(oauth2.NoContext, oauth2.HTTPClient, client)

	token, err := authConfig.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(token)
	// fesp, err := h.client.Get("https://oauth.reddit.com/api/v1/me.json")
	// content, err := ioutil.ReadAll(resp.Body)
	return c.String(http.StatusOK, "hi")
}
