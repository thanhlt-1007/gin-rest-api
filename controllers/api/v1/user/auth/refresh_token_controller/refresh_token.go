package refresh_token_controller

import (
	"github.com/gin-gonic/gin"
)

func RefreshToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(
			200,
			gin.H {
				"message": "OK",
			},
		)
	}
}
