package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func getPingHandler() gin.HandlerFunc {
    return func(context *gin.Context) {
        context.JSON(
            http.StatusOK,
            gin.H {
                "message": "pong",
            },
        )
    }
}

func main() {
    engine := gin.Default()

    engine.GET("/ping", getPingHandler())

    engine.Run()
}
