package controllers

import (
	"gin-sample-project/dto"
	"gin-sample-project/services"
	util "gin-sample-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type auth struct {
// 	Username string `valid:"Required; MaxSize(50)"`
// 	Password string `valid:"Required; MaxSize(50)"`
// }

// @Summary Get Auth
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func PostAuth(c *gin.Context) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Recovered from panic:", r)
	// 	}
	// }()

	var payload *dto.Auth
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := util.Validate(payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"payload-validation-error": err.Error()})
		return
	}

	user, err := services.FetchUserByUserName(payload.Username)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	util.APIResponse(c, "Login successfully", http.StatusOK, user)
}
