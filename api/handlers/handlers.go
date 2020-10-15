package handlers

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/oauth2"
)

// Handler is used so that we can dependency inject different services into the database
type Handler struct {
	db *mongo.Database
}

type transp struct {
	config    *oauth2.Config
	userAgent string
}

// RoundTrip sets headers
func (t *transp) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", t.userAgent)
	//fmt.Sprintf("relevant_for_reddit/0.0 (by /u/%s)", viper.GetString("reddit.username")),
	// 	fmt.Sprintf("Basic ")

	req.SetBasicAuth(t.config.ClientID, t.config.ClientSecret)
	// 	"Basic "+viper.GetString("reddit.client")+":"+viper.GetString("reddit.Secret")+")",
	// )

	return http.DefaultTransport.RoundTrip(req)
}

// NewHandler with given database and other dependencies
func NewHandler(db *mongo.Database) *Handler {
	return &Handler{
		db,
	}
}
