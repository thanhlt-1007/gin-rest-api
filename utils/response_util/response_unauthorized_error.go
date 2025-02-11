package response_util

import (
	"net/http"

	"gin-rest-api/errors"
	"gin-rest-api/serializers/errors/unauthorized_error_serializer"
	"github.com/gin-gonic/gin"
)

func ResponseUnauthorizedError(context *gin.Context, err errors.UnauthorizedError) {
	context.JSON(
		http.StatusUnauthorized,
		unauthorized_error_serializer.Serializer(err),
	)
}
