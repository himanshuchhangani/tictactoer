// handlers/handlers_test.go

package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetBoard(t *testing.T) {
	// Create a new Gin router
	router := setupRouter()

	// Create a test request to the /board endpoint
	req, _ := http.NewRequest("GET", "/board", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rec.Code)

	// Parse the response JSON
	var responseBody map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &responseBody)
	assert.Nil(t, err)

	// Verify expected keys in the response
	assert.Contains(t, responseBody, "currentTurn")
	assert.Contains(t, responseBody, "winner")
	assert.Contains(t, responseBody, "board")
}

func TestMove(t *testing.T) {
	// Create a new Gin router
	router := setupRouter()

	// Define the request payload
	requestPayload := `{"player":"X","row":0,"column":0}`

	// Create a test request to the /game/move endpoint with the payload
	req, _ := http.NewRequest("POST", "/game/move", strings.NewReader(requestPayload))
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rec.Code)

	// Parse the response JSON
	var responseBody map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &responseBody)
	assert.Nil(t, err)

	// Verify expected keys in the response
	assert.Contains(t, responseBody, "currentTurn")
	assert.Contains(t, responseBody, "winner")
	assert.Contains(t, responseBody, "board")
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// Add your routes and handlers here
	router.GET("/board", GetBoard)
	router.POST("/game/move", Move)
	router.DELETE("/game/reset", Delete)

	return router
}
