package response_util

import (
	"net/http"

	"gin-rest-api/serializers/errors/internal_server_error_serializer"
	"github.com/gin-gonic/gin"
)

func ResponseInternalServerError(context *gin.Context, err error) {
	context.JSON(
		http.StatusInternalServerError,
		internal_server_error_serializer.Serializer(err),
	)
}
