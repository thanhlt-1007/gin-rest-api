package router

import (
	"log"

	"gin-rest-api/controllers/api/v1/user/auth/refresh_token_controller"
	"gin-rest-api/controllers/api/v1/user/auth/sign_in_controller"
	"gin-rest-api/controllers/api/v1/user/auth/sign_up_controller"
	"gin-rest-api/controllers/api/v1/user/me_controller"
	"gin-rest-api/controllers/ping_controller"
	"gin-rest-api/initializers"
	"gin-rest-api/middlewares/api/v1/user_middleware"
)

func Init() {
	log.Printf("Init router")
	ping()
	apiV1UserAuth()
	apiV1User()
}

func ping() {
	initializers.ENGINE.GET("/ping", ping_controller.GetPing())
}

func apiV1UserAuth() {
	authRouterGroup := initializers.ENGINE.Group("/api/v1/user/auth")
	authRouterGroup.POST("/sign_up", sign_up_controller.SignUp())
	authRouterGroup.POST("/sign_in", sign_in_controller.SignIn())
	authRouterGroup.POST("/refresh_token", refresh_token_controller.RefreshToken())
}

func apiV1User() {
	authRouterGroup := initializers.ENGINE.Group("/api/v1/user", user_middleware.Authenticate())
	authRouterGroup.GET("/me", me_controller.GetMe())
}
