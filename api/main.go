package main

import (
	"fmt"

	"github.com/ablades/relevant/server"
	"github.com/spf13/viper"
)

func main() {

	// Set the file name of the configurations file
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Error reading config file, %s", err))
	}

	// Start Server
	server := server.NewServer(nil)

	server.Start(":8000")
}
