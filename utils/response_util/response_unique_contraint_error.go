package response_util

import (
	"net/http"

	"gin-rest-api/errors"
	"gin-rest-api/serializers/errors/unique_contraint_error_serializer"
	"github.com/gin-gonic/gin"
)

func ResponseUniqueContraintError(context *gin.Context, err errors.UniqueContraintError) {
	context.JSON(
		http.StatusUnprocessableEntity,
		unique_contraint_error_serializer.Serializer(err),
	)
}
