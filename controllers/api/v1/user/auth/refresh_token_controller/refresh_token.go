package refresh_token_controller

import (
	_ "gin-rest-api/responses/api/v1/user/auth/refresh_token_controller/refresh_token_response"
	_ "gin-rest-api/responses/errors/internal_server_error_response"
	_ "gin-rest-api/responses/errors/unauthorized_error_response"
	"gin-rest-api/serializers/api/v1/user/auth/refresh_token_controller/refresh_token_serializer"
	"gin-rest-api/services/api/v1/user/auth/refresh_token_controller/refresh_token_service"
	"gin-rest-api/utils/response_util"

	"github.com/gin-gonic/gin"
)

// godoc
// @Tags User-Auth
// @Router /user/auth/refresh_token [post]
// @Summary User refresh token
// @Accept json
// @Produce json
// @Success 200	{object} refresh_token_response.Response
// @Failure 401 {object} unauthorized_error_response.Response
// @Failure 500 {object} internal_server_error_response.Response
func RefreshToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		refresh_token := context.GetHeader("X-Refresh-Token")

		token := refresh_token_service.Perform(refresh_token)

		response_util.ResponseOK(
			context,
			refresh_token_serializer.Serializer(token),
		)
	}
}
