package main

import (
	"os"
	"twc/routes"
)

// main is the entry point of the application
func main() {
	// Set up the routes using the router defined in the routes package
	router := routes.SetupRoutes()

	// Parse the environment variable for the server port or use the default value
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080" // Default port
	}

	// Start the Gin server using the configured router and port
	startServer(router, port)
}
