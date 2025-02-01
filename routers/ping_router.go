package routers

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

func SetPingRouters(engine *gin.Engine) {
    engine.GET("/ping", getPingHandler())
}
