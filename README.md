Tic-Tac-Toe Game Server

A simple web-based Tic-Tac-Toe game server built with Go and the Gin framework. This project provides HTTP endpoints to play and manage a game of Tic-Tac-Toe. The server features include getting the game board, making moves, and resetting the game.

## Features

- RESTful API endpoints for playing Tic-Tac-Toe
- Get the current game board, including the current player's turn and winner status
- Make moves in the game by specifying the player, row, and column
- Reset the game to its initial state

## Technologies

- Go programming language
- Gin web framework
- Docker for containerization

## Getting Started

1. Clone the repository: `git clone https://github.com/himanshuchhangani/tictactoer.git`
2. Install dependencies: `go mod tidy`
3. Run the server: `go run main.go`


## Dockerization 

To run the server in a Docker container:

Build the Docker image:

      docker build -t tic-tac-toe-server .

Run the Docker container:

      docker run -p 8080:8080 tic-tac-toe-server


## Playing the Game

GET /board: Get the current game board, including the current player's turn and the winner (if any).

POST /game/move: Make a move in the game. Send a JSON payload with the player making the move, the row, and the column.

DELETE /game/reset: Reset the game to its initial state.

sample request : 

  
     curl -X POST -H "Content-Type: application/json" -d '{
    "player": "X",
    "row": 0,
    "column": 0
    }' http://localhost:8080/game/move

## Testing

To run tests, execute the following command:
    
    go test ./...



## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

