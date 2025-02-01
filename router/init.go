package router

import (
    "gin-rest-api/controllers/api/v1/user/auth_controller"
    "gin-rest-api/controllers/ping_controller"
    "gin-rest-api/initializers"
)

func Init() {
    ping()
    apiV1UserAuth()
}

func ping() {
    initializers.ENGINE.GET("/ping", ping_controller.GetPing())
}

func apiV1UserAuth() {
    authRouterGroup := initializers.ENGINE.Group("/api/v1/user/auth")
    authRouterGroup.POST("/sign-up", auth_controller.SignUp())
}
