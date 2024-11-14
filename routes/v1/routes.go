package v1

import (
	"gin-sample-project/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutesV1 sets up routes for version 1

func RegisterRoutesV1(router *gin.RouterGroup) {

	user := router.Group("/user")
	{
		user.GET("/:id", controllers.GetUser)
		user.POST("/", controllers.CreateUser)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/", controllers.PostAuth)

	}
}
