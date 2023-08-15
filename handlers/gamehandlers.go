package handlers

import (
	"net/http"

	"twc/models"

	ttt "github.com/chrisfregly/tictactoe"
	"github.com/gin-gonic/gin"
)

// MoveRequest represents the JSON request structure for a move
type MoveRequest struct {
	Player ttt.Player `json:"player"`
	Row    int        `json:"row"`
	Column int        `json:"column"`
}

// GetBoard handles the GET request to retrieve the current game board status
func GetBoard(c *gin.Context) {
	// Format the response and send it as JSON
	response := formatResponse()
	c.JSON(http.StatusOK, response)
}

// Delete handles the DELETE request to reset the game state
func Delete(c *gin.Context) {
	// Reset the game state using the models package
	models.Reset()
	c.JSON(http.StatusOK, "Current Game Deleted Successfully.")
}

// Move handles the POST request to make a move on the game board
func Move(c *gin.Context) {
	var moveReq MoveRequest

	// Parse the JSON request body into a MoveRequest struct
	if err := c.BindJSON(&moveReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Convert player's string to Player type
	player := ttt.Player(moveReq.Player)

	// Make a move on the game board
	err := models.Game.Move(player, moveReq.Row, moveReq.Column)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Format the response and send it as JSON
	response := formatResponse()
	c.JSON(http.StatusOK, response)
}

// formatResponse formats the game state into a response structure
func formatResponse() gin.H {
	// Get the current turn from the global Game instance
	currentTurn := string(models.Game.GetTurn())

	// Determine the winner (if any)
	var winner *string
	if models.Game.GetWinner() != nil {
		winnerName := string(*models.Game.GetWinner())
		winner = &winnerName
	}

	// Create a representation of the game board
	var board [ttt.BoardSize][ttt.BoardSize]*string
	for i, row := range models.Game.GetBoard() {
		for j, player := range row {
			if player != nil {
				playerName := string(*player)
				board[i][j] = &playerName
			} else {
				empty := " "
				board[i][j] = &empty
			}
		}
	}

	// Construct the response structure
	return gin.H{
		"currentTurn": currentTurn,
		"winner":      winner,
		"board":       board,
	}
}
