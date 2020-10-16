package handlers

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/ablades/relevant/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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

//Send Requests to Reddit api
func (h *Handler) request(ctx echo.Context, method string, url string, body io.Reader) []byte {

	// Build Request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Error(err)
	}
	req.Header.Set("User-Agent", viper.GetString("reddit.agent"))

	// Send Request
	res, err := h.client.Do(req)
	if err != nil {
		log.Error(err)
	}
	defer res.Body.Close()

	// Read response
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
	}

	return content
}
