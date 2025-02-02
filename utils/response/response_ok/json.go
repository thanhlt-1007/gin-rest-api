package response_ok

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func JSON(context *gin.Context, data any) {
    context.JSON(
        http.StatusOK,
        gin.H {
            "data": data,
        },
    )
}
