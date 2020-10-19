package handlers

import (
	"fmt"
	"net/http"

	"github.com/ablades/relevant/models"
	"github.com/labstack/echo/v4"
)

// UserHome page
func (h *Handler) UserHome(c echo.Context) (err error) {
	userName := c.Param("name")
	fmt.Println(userName)
	uStore := models.GetUserStore(h.db)
	user := uStore.GetUserByName(userName)

	return c.JSON(http.StatusOK, user)
}

// UpdateUserSubscriptions for active user
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
	c.SetCookie(&http.Cookie{
		Name:  "username",
		Value: userName,
		Path:  "/",
	})
	// create user
	if user == nil {
		// Add list of subreddits to a user objects subs
		subreddits := h.getRedditUserSubs()
		subs := make(map[string][]string)
		for subName := range subreddits {
			subs[subName] = make([]string, 0)
		}
		// Create user
		newUser := models.User{
			Name: userName,
			Subs: subs,
		}
		// Insert user into db
		uStore.CreateUser(newUser)
	}

	return c.Redirect(http.StatusPermanentRedirect, "http://localhost:3000/user/"+userName)
}

// DeleteUser and remove existing content
func (h *Handler) DeleteUser(c echo.Context) (err error) {
	userName := c.Param("name")
	uStore := models.GetUserStore(h.db)
	uStore.DeleteUserByName(userName)

	return c.String(http.StatusGone, "Deleted:"+userName)
}

// GetUserSubBanners from DB
func (h *Handler) GetUserSubBanners(c echo.Context) (err error) {
	uStore := models.GetUserStore(h.db)
	subStore := models.GetSubRedditStore(h.db)
	userName := c.Param("name")

	user := uStore.GetUserByName(userName)
	banners := make(map[string]string)

	for subName := range user.Subs {
		banners[subName] = subStore.GetSubReddit(subName).BannerUrl
	}

	return c.JSON(http.StatusOK, banners)
}

// func getUser(c echo.Context) error {
// 	return
// }

// func updateUser(c echo.Context) error {
// }

// func deleteUser(c echo.Context) error {

// }
