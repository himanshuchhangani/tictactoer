// server.go

package main

import (
	"github.com/gin-gonic/gin"
)

// startServer starts the Gin server on the specified port
func startServer(router *gin.Engine, port string) {
	// Start the Gin server using the provided router and port
	router.Run(port)
}
