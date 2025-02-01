package ping_controller

import (
    "gin-rest-api/responses"
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetPing() gin.HandlerFunc {
    return func(context *gin.Context) {
        context.JSON(
            http.StatusOK,
            responses.PingResponse.GetPingResponse,
        )
    }
}
