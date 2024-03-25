package main

import (
	"post/server"

	"github.com/GiveGetGo/shared/config"
)

func main() {
	config.LoadEnv(".env.post") // Load environment variables

	server.StartServer() // Start the server
}
