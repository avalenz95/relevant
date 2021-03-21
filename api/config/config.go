package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

// GetAuthConfig for app
func GetAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("REDDIT_APP"),
		ClientSecret: os.Getenv("REDDIT_SECRET"),
		Scopes:       []string{"mysubreddits", "identity", "history"},
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://www.reddit.com/api/v1/access_token",
			AuthURL:  "https://reddit.com/api/v1/authorize",
		},
		RedirectURL: os.Getenv("REDDIT_REDIRECT"),
	}
}

func LoadENV() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
