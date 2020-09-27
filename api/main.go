package main

import (
	"github.com/ablades/relevant/config"
	"github.com/ablades/relevant/server"
)

func main() {
	config.Init()
	// Start Server
	server := server.NewServer(nil)

	server.Start(":8000")
}
