package main

import (
	"github.com/ablades/relevant/server"
)

func main() {

	// Start Server
	server := server.NewServer(nil)
	server.Start(":8000")
}
