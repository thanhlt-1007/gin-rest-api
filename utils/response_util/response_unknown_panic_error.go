package response_util

import (
	"net/http"

	"gin-rest-api/serializers/errors/unknown_panic_error_serializer"

	"github.com/gin-gonic/gin"
)

func ResponseUnknownPanicError(context *gin.Context, err any) {
	context.JSON(
		http.StatusInternalServerError,
		unknown_panic_error_serializer.Serializer(err),
	)
}
