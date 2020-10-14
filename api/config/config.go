package config

import (
	"fmt"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

// GetAuthConfig for app
func GetAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     viper.GetString("reddit.client"),
		ClientSecret: viper.GetString("reddit.secret"),
		Scopes:       []string{"mysubreddits", "identity", "history"},
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://www.reddit.com/api/v1/access_token",
			AuthURL:  "https://reddit.com/api/v1/authorize",
		},
		RedirectURL: viper.GetString("reddit.redirect"),
	}
}

func init() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config file, %s", err))
	}
}
