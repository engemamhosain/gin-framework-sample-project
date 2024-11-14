package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-sample-project/db"
	"gin-sample-project/dto"
	"gin-sample-project/models"
	"gin-sample-project/routes"
	util "gin-sample-project/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var router = routes.SetupRouter()

func init() {
	db.InitializeMySQL()

	// defer db.DB.Close()

}
func TestGetUser(t *testing.T) {
	// Initialize the router using setupRouter

	// Create a new HTTP request
	req, _ := http.NewRequest("GET", "/api/v1/user/1", nil)

	// Record the HTTP response using httptest
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Body.Len() == 0 {
		t.Fatal("Received empty response body")
	}

	// Check if the status code is 200 (OK)
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body
	var user models.User
	err := json.Unmarshal(w.Body.Bytes(), &user)
	fmt.Println("body", user)
	assert.NoError(t, err)

	// Verify response data
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "john_doe", user.Username)

}

func TestPostUser(t *testing.T) {
	// Initialize the router using setupRouter
	// router := routes.SetupRouter()
	jsonResponse := util.Responses{}

	user := dto.User{
		Username: "user",
		Password: "password",
		Email:    "email",
	}

	// Convert user struct to JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("Error marshalling JSON: %v", err)
	}

	// Create a new HTTP request
	req, _ := http.NewRequest("POST", "/api/v1/user/", bytes.NewBuffer(jsonData))

	// Record the HTTP response using httptest
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Body.Len() == 0 {
		t.Fatal("Received empty response body")
	}

	// Check if the status code is 200 (OK)
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body

	err = json.Unmarshal(w.Body.Bytes(), &jsonResponse)
	fmt.Println("body", jsonResponse)
	assert.NoError(t, err)

	assert.Equal(t, "Login successfully", jsonResponse.Message)

}
