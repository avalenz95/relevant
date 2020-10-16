package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

// AuthReddit route
func (h *Handler) AuthReddit(c echo.Context) (err error) {
	// Set state and additional params
	state := c.Get("csrf").(string)

	url := h.config.AuthCodeURL(
		state,
		oauth2.SetAuthURLParam("response_type", "code"),
		oauth2.SetAuthURLParam("duration", "temporary"),
	)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}
