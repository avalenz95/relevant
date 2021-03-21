package handlers

import (
	"net/http"
	"os"

	"github.com/ablades/relevant/api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/oauth2"
)

// Handler for routes
type Handler struct {
	db     *mongo.Database
	config *oauth2.Config
	client *http.Client
	token  *oauth2.Token
	T      http.RoundTripper
}

// NewHandler with given database and other dependencies
func NewHandler(db *mongo.Database) *Handler {
	return &Handler{
		db,
		config.GetAuthConfig(),
		http.DefaultClient,
		nil,
		http.DefaultTransport,
	}
}

// RoundTrip Sets Headers
func (h *Handler) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", os.Getenv("REDDIT_USER_AGENT"))
	req.SetBasicAuth(h.config.ClientID, h.config.ClientSecret)

	return h.T.RoundTrip(req)
}
