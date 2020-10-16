package handlers

import (
	"net/http"

	"github.com/ablades/relevant/config"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/oauth2"
)

// Handler for routes
type Handler struct {
	db     *mongo.Database
	config *oauth2.Config
	client *http.Client
	token  *oauth2.Token
}

// NewHandler with given database and other dependencies
func NewHandler(db *mongo.Database) *Handler {
	return &Handler{
		db,
		config.GetAuthConfig(),
		http.DefaultClient,
		nil,
	}
}

// RoundTrip Sets Headers
func (h *Handler) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", viper.GetString("reddit.agent"))
	req.SetBasicAuth(h.config.ClientID, h.config.ClientSecret)

	return http.DefaultTransport.RoundTrip(req)
}
