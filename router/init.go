package router

import (
	"log"

	"gin-rest-api/controllers/api/v1/user/auth/sign_up_controller"
	"gin-rest-api/controllers/ping_controller"
	"gin-rest-api/initializers"
)

func Init() {
	log.Printf("Init router")
	ping()
	apiV1UserAuth()
}

func ping() {
	initializers.ENGINE.GET("/ping", ping_controller.GetPing())
}

func apiV1UserAuth() {
	authRouterGroup := initializers.ENGINE.Group("/api/v1/user/auth")
	authRouterGroup.POST("/sign_up", sign_up_controller.SignUp())
}
