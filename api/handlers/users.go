package handlers

import (
	"net/http"

	"github.com/ablades/relevant/models"
	"github.com/labstack/echo/v4"
)

// UserHome page
func (h *Handler) UserHome(c echo.Context) (err error) {
	return
}

// UpdateSubs for active user
func (h *Handler) UpdateUserSubscriptions(c echo.Context) (err error) {
	userName := c.Param("name")
	uStore := models.GetUserStore(h.db)

	subList := h.getRedditUserSubs()

	if uStore.UpdateUserSubs(userName, subList) {
		return c.JSON(http.StatusOK, userName)
	}

	return c.JSON(http.StatusNotModified, subList)
}

// CreateUser add to DB
func (h *Handler) CreateUser(c echo.Context) (err error) {
	userName := c.Param("name")
	uStore := models.GetUserStore(h.db)
	user := uStore.GetUserByName(userName)

	// user already exists
	if user != nil {
		return c.JSON(http.StatusSeeOther, user.Name)
	}
	//userName := h.getRedditUserName()
	// Add list of subreddits to a user objects subs
	subreddits := h.getRedditUserSubs()
	subs := make(map[string][]string)
	for _, subreddit := range subreddits {
		subs[subreddit] = make([]string, 0)
	}
	// Create user
	newUser := models.User{
		Name: userName,
		Subs: subs,
	}

	uStore.CreateUser(newUser)
	//Insert user into db
	return c.JSON(http.StatusCreated, user.Name)
}

// DeleteUser and remove existing content
func (h *Handler) DeleteUser(c echo.Context) (err error) {
	userName := c.Param("name")
	uStore := models.GetUserStore(h.db)
	uStore.DeleteUserByName(userName)

	return c.String(http.StatusGone, "Deleted:"+userName)
}

// func getUser(c echo.Context) error {
// 	return
// }

// func updateUser(c echo.Context) error {
// }

// func deleteUser(c echo.Context) error {

// }
