package handlers

import (
	"fmt"
	"net/http"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

// Handler is used so that we can dependency inject different services into the database
type Handler struct {
	db     *mongo.Database
	client *http.Client
	tp     http.RoundTripper
}

// RoundTrip sets headers
func (h *Handler) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set(
		"User-Agent",
		"relevant_for_reddit/1.0 (by /u/"+viper.GetString("reddit.username")+")",
	)

	req.Header.Set(
		"Authorization",
		"Basic "+viper.GetString("reddit.client")+":"+viper.GetString("reddit.Secret")+")",
	)
	fmt.Println("HERE")
	return http.DefaultTransport.RoundTrip(req)
}

// NewHandler with given database and other dependencies
func NewHandler(db *mongo.Database) *Handler {
	return &Handler{
		db,
		http.DefaultClient,
		http.DefaultTransport,
	}
}
