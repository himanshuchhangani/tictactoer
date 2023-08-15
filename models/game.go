package models

import ttt "github.com/chrisfregly/tictactoe"

// Game represents the global instance of the TicTacToe game
var Game ttt.TicTacToe

// init initializes the Game instance with a new TicTacToe game when the package is imported
func init() {
	// Initialize the global Game instance with a new TicTacToe game
	Game = ttt.NewTicTacToe()
}

// Reset resets the game state by creating a new TicTacToe game instance
func Reset() {
	// Reset the global Game instance by creating a new TicTacToe game
	Game = ttt.NewTicTacToe()
}
