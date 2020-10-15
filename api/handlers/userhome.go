package handlers

import "github.com/labstack/echo/v4"

// UserHome page
func (h *Handler) UserHome(c echo.Context) (err error) {

	h.verifySession(c)
	return
}
