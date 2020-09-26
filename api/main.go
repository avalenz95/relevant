package main

import (
	"fmt"
	"os"

	"github.com/ablades/relevant/server"
)

func main() {
	fmt.Println(os.Getenv("DB_TEST_USERNAME"))
	// Start Server
	server := server.NewServer(nil)

	server.Start(":8000")
}
