package response_unique_contraint_error

import (
    "gin-rest-api/errors"
    "gin-rest-api/serializers/errors/unique_contraint_error_serializer"
    "github.com/gin-gonic/gin"
    "net/http"
)

func JSON(context *gin.Context, err errors.UniqueContraintError) {
    context.JSON(
        http.StatusBadRequest,
        unique_contraint_error_serializer.Serializer(err),
    )
}
