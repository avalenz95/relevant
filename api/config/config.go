package config

import (
	"os"

	"github.com/joho/godotenv"
)

// DBConfig Holds DB Configuration Info
type DBConfig struct {
	URI      string
	Username string
	Password string
}

func init() {
	godotenv.Load()
}

// GetDBConfig will get the database config values from the envrionment. They will be loaded in the init func in config.go
func GetDBConfig() *DBConfig {
	return &DBConfig{
		URI:      os.Getenv("DB_URI"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}

// GetTestDBConfig will return a test database's config values if they are set and you want to create a test database
func GetTestDBConfig() *DBConfig {
	return &DBConfig{
		URI:      os.Getenv("DB_TEST_URI"),
		Username: os.Getenv("DB_TEST_USERNAME"),
		Password: os.Getenv("DB_TEST_PASSWORD"),
	}
}
