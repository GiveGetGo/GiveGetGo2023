package main

import (
	"match/server"

	"github.com/GiveGetGo/shared/config"
)

func main() {
	config.LoadEnv(".env.match") // Load environment variables

	server.StartServer() // Start the server
}
