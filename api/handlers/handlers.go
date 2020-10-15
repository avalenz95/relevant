package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/ablades/relevant/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func (h *Handler) verifySession(c echo.Context) error {
	cookie, err := c.Cookie("_csrf")
	if err != nil {
		log.Error(err)
	}

	// Verify csrfs match
	if cookie.Value != c.Get(middleware.DefaultCSRFConfig.ContextKey) {
		return c.String(http.StatusForbidden, "You do not have access to this page")
	}

	return c.String(http.StatusOK, "Valid")
}

//Send Requests to Reddit api
func (h *Handler) request(method string, url string) []byte {

	// Build Request
	req, err := http.NewRequest(method, url, nil)
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
