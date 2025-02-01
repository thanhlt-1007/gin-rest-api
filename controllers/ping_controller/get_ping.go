package ping_controller

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetPing() gin.HandlerFunc {
    return func(context *gin.Context) {
        context.JSON(
            http.StatusOK,
            gin.H {
                "message": "pong",
            },
        )
    }
}
