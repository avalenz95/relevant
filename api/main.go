package main

import (
	"os"

	"github.com/ablades/relevant/api/config"
	"github.com/ablades/relevant/api/server"
)

func main() {

	config.LoadENV()
	// Start Server
	server := server.NewServer(nil)

	server.Start(os.Getenv("PORT"))
}
