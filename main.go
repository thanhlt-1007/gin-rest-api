package main

import (
	"log"

	"gin-rest-api/database"
	"gin-rest-api/initializers"
	"gin-rest-api/middlewares"
	"gin-rest-api/router"
	"gin-rest-api/swagger"
	"gin-rest-api/utils/fmt_util"
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
	err := initializers.ENGINE.Run()
	if err != nil {
		log.Printf("initializers.ENGINE.Run() error\n%s", fmt_util.SprintfError(err))
		panic(err)
	}
}
