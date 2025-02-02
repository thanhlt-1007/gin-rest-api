package response

import (
    "gin-rest-api/serializers/errors/unknown_panic_error_serializer"
    "github.com/gin-gonic/gin"
    "net/http"
)

func ResponseUnkownPanicError(context *gin.Context, err any) {
    context.JSON(
        http.StatusInternalServerError,
        unknown_panic_error_serializer.Serializer(err),
    )
}
