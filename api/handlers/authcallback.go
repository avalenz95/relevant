package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ablades/relevant/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

// AuthCallback for reddit authentication
func (h *Handler) AuthCallback(c echo.Context) (err error) {
	c.Request().Header.Set(
		"User-Agent",
		"relevant_for_reddit/1.0 (by /u/"+viper.GetString("reddit.username")+")",
	)
	// Get Query Parameters
	code := c.QueryParam("code")
	queryState := c.QueryParam("state")
	state := c.Get(middleware.DefaultCSRFConfig.ContextKey)
	// set context
	ctx := c.Request().Context()
	c.Request().WithContext()
	c.
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
	h.client = authConfig.Client(ctx, token)
	// TODO: Stuck on client
	resp, err := h.client.Get("https://oauth.reddit.com/api/v1/me")
	fmt.Println(resp.Header.Get("User-Agent"))
	content, err := ioutil.ReadAll(resp.Body)

	return c.JSONBlob(http.StatusOK, content)
}
