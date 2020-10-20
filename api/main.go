package main

import (
	"github.com/ablades/relevant/daemon"
	"github.com/ablades/relevant/server"
)

func main() {

	// Start Server
	server := server.NewServer(nil)
	daemon.Run()
	server.Start(":8000")
}
