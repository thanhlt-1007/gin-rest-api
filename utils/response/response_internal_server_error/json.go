package response_internal_server_error

import (
    "gin-rest-api/serializers/errors/internal_server_error_serializer"
    "github.com/gin-gonic/gin"
    "net/http"
)

func JSON(context *gin.Context, err error) {
    context.JSON(
        http.StatusInternalServerError,
        internal_server_error_serializer.Serializer(err),
    )
}
