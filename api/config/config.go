package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Init Env Variables
func init() {
	godotenv.Load()
}

// DBConfig data
type DBConfig struct {
	URI      string
	Username string
	Password string
}

// RedditConfig data
type RedditConfig struct {
	Username string
	Password string
	Client   string
	Secret   string
}

// GetRedditConfig data from env
func GetRedditConfig() *RedditConfig {
	return &RedditConfig{
		Username: os.Getenv("REDDIT_USERNAME"),
		Password: os.Getenv("REDDIT_PASSWORD"),
		Client:   os.Getenv("TEST_APP_CLIENT"),
		Secret:   os.Getenv("TEST_APP_SECRET"),
	}
}

// GetDBConfig data from env
func GetDBConfig() *DBConfig {
	return &DBConfig{
		URI:      os.Getenv("DB_TEST_URI"),
		Username: os.Getenv("DB_TEST_USERNAME"),
		Password: os.Getenv("DB_TEST_PASSWORD"),
	}
}
