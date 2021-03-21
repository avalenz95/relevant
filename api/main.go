package main

import (
	"github.com/ablades/relevant/api/config"
	"github.com/ablades/relevant/api/server"
)

func main() {

	config.LoadENV()
	// Start Server
	server := server.NewServer(nil)

	server.Start(":8000")
}
