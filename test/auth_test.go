package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-sample-project/db"
	"gin-sample-project/dto"
	"gin-sample-project/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	db.InitializeMySQL()
}
func TestPostAuth(t *testing.T) {
	// Initialize the router using setupRouter
	router := routes.SetupRouter()

	auth := dto.Auth{
		Username: "john_doe",
		Password: "test",
	}

	// Convert user struct to JSON
	jsonData, err := json.Marshal(auth)
	if err != nil {
		t.Fatalf("Error marshalling JSON: %v", err)
	}

	// Create a new HTTP request
	req, _ := http.NewRequest("POST", "/api/v1/auth/", bytes.NewBuffer(jsonData))

	// Record the HTTP response using httptest
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Body.Len() == 0 {
		t.Fatal("Received empty response body")
	}

	// Check if the status code is 200 (OK)
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body

	err = json.Unmarshal(w.Body.Bytes(), &auth)
	fmt.Println("body", auth)
	assert.NoError(t, err)

	assert.Equal(t, "john_doe", auth.Username)

}
