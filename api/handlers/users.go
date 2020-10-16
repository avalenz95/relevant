package handlers

import (
	"net/http"

	"github.com/ablades/relevant/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserHome page
func (h *Handler) UserHome(c echo.Context) (err error) {
	return
}

// CreateUser add to DB
func (h *Handler) CreateUser(c echo.Context) error {

	uStore := models.GetUserStore(h.db)

	user := models.User{
		ID: primitive.NewObjectID(),
	}
	// Bind Payload to model
	err := c.Bind(user)
	if err != nil {
		log.Error(err)
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
