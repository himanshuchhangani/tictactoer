// routes.go

package routes

import (
	"twc/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the routes and handlers for the TicTacToe game
func SetupRoutes() *gin.Engine {
	// Create a new Gin router with default middleware
	router := gin.Default()

	// Register the GET endpoint for retrieving the game board
	router.GET("/game", handlers.GetBoard)

	// Register the POST endpoint for making a move on the game board
	router.POST("/game/move", handlers.Move)

	// Register the DELETE endpoint for resetting the game state
	router.DELETE("/game", handlers.Delete)

	// Return the configured router
	return router
}
