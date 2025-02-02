package response

import (
    "gin-rest-api/serializers/errors/validation_error_serializer"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "net/http"
)

func ResponseValidationError(context *gin.Context, err validator.ValidationErrors) {
    context.JSON(
        http.StatusBadRequest,
        validation_error_serializer.Serializer(err),
    )
}
