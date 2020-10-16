package handlers

import (
	"net/http"

	"github.com/ablades/relevant/models"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserHome page
func (h *Handler) UserHome(c echo.Context) (err error) {
	return
}

// CreateUser add to DB
func (h *Handler) CreateUser(c echo.Context) error {
	// Get User Info from endpoint

	userName := h.getRedditUserName()
	// Add list of subreddits to a user objects subs
	subreddits := h.getRedditUserSubs()
	subs := make(map[string][]string)
	for _, subreddit := range subreddits {
		subs[subreddit] = make([]string, 0)
	}
	// Create user
	uStore := models.GetUserStore(h.db)

	user := models.User{
		ID:   primitive.NewObjectID(),
		Name: userName,
		Subs: subs,
	}

	uStore.CreateUser(user)
	//Insert user into db
	return c.JSON(http.StatusCreated, user)
}

// func getUser(c echo.Context) error {
// 	return
// }

// func updateUser(c echo.Context) error {
// }

// func deleteUser(c echo.Context) error {

// }
