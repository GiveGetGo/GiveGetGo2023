package main

import (
	"bid/server"

	"github.com/GiveGetGo/shared/config"
)

func main() {
	config.LoadEnv(".env.bid") // Load environment variables

	server.StartServer() // Start the server
}
