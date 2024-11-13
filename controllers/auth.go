package controllers

import (
	util "gin-sample-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary Get Auth
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	// valid := validation.Validation{}

	// username := c.PostForm("username")
	// password := c.PostForm("password")

	// a := auth{Username: username, Password: password}
	// ok, _ := valid.Valid(&a)

	// if !ok {
	// 	//common.MarkErrors(valid.Errors)
	// 	common.Response(http.StatusBadRequest, common.INVALID_PARAMS, nil)
	// 	return
	// }

	util.APIResponse(c, "Login successfully", http.StatusOK, http.MethodPost, map[string]string{"accessToken": "accessToken"})
}
