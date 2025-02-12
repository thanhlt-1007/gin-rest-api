package sign_in_controller

import (
	"gin-rest-api/requests/api/v1/user/auth/sign_in_controller/sign_in_request"
	_ "gin-rest-api/responses/api/v1/user/auth/sign_in_controller/sign_in_response"
	_ "gin-rest-api/responses/errors/internal_server_error_response"
	_ "gin-rest-api/responses/errors/unauthorized_error_response"
	"gin-rest-api/serializers/api/v1/user/auth/sign_in_controller/sign_in_serializer"
	"gin-rest-api/services/api/v1/user/auth/sign_in_controller/sign_in_service"
	"gin-rest-api/utils/response_util"
	"github.com/gin-gonic/gin"
)

// godoc
// @Tags User-Auth
// @Router /user/auth/sign_in [post]
// @Summary User sign in
// @Accept json
// @Produce json
// @Param request body sign_in_request.Request true "Sign in request body"
// @Success 200	{object} sign_in_response.Response
// @Failure 401 {object} unauthorized_error_response.Response
// @Failure 500 {object} internal_server_error_response.Response
func SignIn() gin.HandlerFunc {
	return func(context *gin.Context) {
		var request sign_in_request.Request
		err := context.ShouldBind(&request)
		if err != nil {
			panic(err)
		}

		token := sign_in_service.Perform(request)

		response_util.ResponseOK(
			context,
			sign_in_serializer.Serializer(token),
		)
	}
}
