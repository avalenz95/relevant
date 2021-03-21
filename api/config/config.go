package config

import (
	"os"

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

// func init() {
// 	// Set the file name of the configurations file
// 	viper.SetConfigName("config")
// 	viper.SetConfigType("yml")
// 	viper.AddConfigPath("..")
// 	viper.AutomaticEnv()

// 	if err := viper.ReadInConfig(); err != nil {
// 		panic(fmt.Errorf("Error reading config file, %s", err))
// 	}
// }
