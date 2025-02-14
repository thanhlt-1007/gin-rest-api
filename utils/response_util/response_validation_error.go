package response_util

import (
	"net/http"

	"gin-rest-api/serializers/errors/validation_error_serializer"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ResponseValidationError(context *gin.Context, err validator.ValidationErrors) {
	context.JSON(
		http.StatusUnprocessableEntity,
		validation_error_serializer.Serializer(err),
	)
}
