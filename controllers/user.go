package controllers

import (
	"fmt"
	"gin-sample-project/dto"
	"gin-sample-project/services"
	util "gin-sample-project/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUser handles GET requests to fetch a user by ID
func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	fmt.Println("id", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := services.FetchUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// CreateUser handles POST requests to create a new user
func CreateUser(c *gin.Context) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	var payload *dto.User
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := util.Validate(payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"payload-validation-error": err.Error()})
		return
	}

	if err := services.RegisterUser(payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		fmt.Println(err)
		return
	}

	util.APIResponse(c, "Login successfully", http.StatusOK, map[string]string{"message": "User created successfully"})
	return

}
