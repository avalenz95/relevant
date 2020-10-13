package handlers

import (
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// RedditAuth route
func (h *Handler) RedditAuth(c echo.Context) (err error) {
	//Build Auth URL
	url := url.URL{
		Scheme: "https",
		Path:   "reddit.com/api/v1/authorize",
	}
	//TODO: ADD state redirect_uri and response_type and state
	q := url.Query()
	q.Add("client_id", viper.GetString("reddit.client"))
	q.Add("response_type", "")
	q.Add("state", "")
	q.Add("redirect_uri", "")
	q.Add("duration", "temporary")
	q.Add("scope", "mysubreddits identity history")

	url.Query().Encode()
	// TODO: LOG redirect
	return c.Redirect(http.StatusTemporaryRedirect, url.String())
}
