package response_util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseOK(context *gin.Context, data any) {
	context.JSON(
		http.StatusOK,
		data,
	)
}
