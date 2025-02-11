package response_ok

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSON(context *gin.Context, data any) {
	context.JSON(
		http.StatusOK,
		data,
	)
}
