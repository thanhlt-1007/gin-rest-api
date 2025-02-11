package main

import (
	"log"

	"gin-rest-api/database"
	"gin-rest-api/initializers"
	"gin-rest-api/middlewares"
	"gin-rest-api/router"
	"gin-rest-api/swagger"
	"gin-rest-api/validators"
)

func init() {
	initializers.Init()
	database.Migrate()
	validators.Init()
	middlewares.Init()
	swagger.Init()
	router.Init()
}

func main() {
	log.Printf("Run engine")
	initializers.ENGINE.Run()
}
