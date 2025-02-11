package middlewares

import (
	"log"

	"gin-rest-api/initializers"
)

func Init() {
	log.Println("Init middlewares")
	initializers.ENGINE.Use(Recovery())
}
