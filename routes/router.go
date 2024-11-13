package routes

import (
	"gin-sample-project/config"
	v1 "gin-sample-project/routes/v1"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func SetupRouter() *gin.Engine {
	env := config.Config.GetString("app.environment")
	if env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	apiV1 := router.Group("/api/v1")
	v1.RegisterRoutesV1(apiV1)

	return router
}
