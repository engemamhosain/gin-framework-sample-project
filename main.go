package main

import (
	"gin-sample-project/config"
	"gin-sample-project/db"
	"gin-sample-project/routes"
	"log"
)

func main() {

	db.InitializeMySQL()
	defer db.DB.Close()

	/**
	@description Setup Server
	*/
	router := routes.SetupRouter()
	/**
	@description Run Server
	*/
	log.Fatal(router.Run(":" + config.Config.GetString("app.port")))
}
