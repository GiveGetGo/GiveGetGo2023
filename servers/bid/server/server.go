package server

import (
	"bid/db"
	"bid/middleware"
)

func StartServer() {
	DB := db.InitDB()                      // Initialize the database
	redisClient := middleware.SetupRedis() // Set up Redis

	r := NewRouter(DB, redisClient) // Set up the router and v1 routes
	r.Run(":8080")                  // Start the server
}
