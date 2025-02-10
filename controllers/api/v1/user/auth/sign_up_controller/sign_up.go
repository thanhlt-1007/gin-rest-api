package sign_up_controller

import (
    _ "gin-rest-api/responses/api/v1/user/auth/sign_up_controller/sign_up_response"
    _ "gin-rest-api/responses/errors/internal_server_error_response"
    _ "gin-rest-api/responses/errors/validation_error_response"
    "gin-rest-api/requests/api/v1/user/auth/sign_up_controller/sign_up_request"
    "gin-rest-api/serializers/api/v1/user/auth/sign_up_controller/sign_up_serializer"
    "gin-rest-api/services/api/v1/user/auth/sign_up_controller/sign_up_service"
    "gin-rest-api/utils/response/response_ok"
    "github.com/gin-gonic/gin"
)

// godoc
// @Tags User-Auth-SignUp
// @Router /user/auth/sign_up [post]
// @Summary User sign up
// @Accept json
// @Produce json
// @Param request body sign_up_request.Request true "Sign up request body"
// @Success 200	{object} sign_up_response.Response
// @Failure 422 {object} validation_error_response.Response
// @Failure 500 {object} internal_server_error_response.Response
func SignUp() gin.HandlerFunc {
    return func(context *gin.Context) {
        var request sign_up_request.Request
        err := context.ShouldBind(&request)
        if err != nil {
            panic(err)
        }

        token := sign_up_service.Perform(request)

        response_ok.JSON(
            context,
            sign_up_serializer.Serializer(token),
        )
    }
}
