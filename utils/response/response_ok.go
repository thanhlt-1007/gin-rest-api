package response

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func ResponseOK(context *gin.Context, data any) {
    context.JSON(
        http.StatusOK,
        gin.H {
            "data": data,
        },
    )
}
