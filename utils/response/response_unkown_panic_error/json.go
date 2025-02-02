package response_unkown_panic_error

import (
    "gin-rest-api/serializers/errors/unknown_panic_error_serializer"
    "github.com/gin-gonic/gin"
    "net/http"
)

func JSON(context *gin.Context, err any) {
    context.JSON(
        http.StatusInternalServerError,
        unknown_panic_error_serializer.Serializer(err),
    )
}
