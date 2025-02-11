package me_controller

import (
	"net/http"

	"gin-rest-api/models"

	"gin-rest-api/consts/middlewares/api/v1/user_middleware/authenticate_consts"
	_ "gin-rest-api/responses/api/v1/user/me_controller/get_me_response"
	_ "gin-rest-api/responses/errors/internal_server_error_response"
	_ "gin-rest-api/responses/errors/unauthorized_error_response"
	"gin-rest-api/serializers/api/v1/user/me_controller/get_me_serializer"
	"github.com/gin-gonic/gin"
)

// godoc
// @Tags User-Me
// @Router /user/me [get]
// @Summary User get info
// @Accept json
// @Produce json
// @Success 200	{object} get_me_response.Response
// @Failure 401 {object} unauthorized_error_response.Response
// @Failure 500 {object} internal_server_error_response.Response
func GetMe() gin.HandlerFunc {
	return func(context *gin.Context) {
		value, _ := context.Get(authenticate_consts.AUTHENTICATED_USER)
		user, _ := value.(*models.User)

		context.JSON(
			http.StatusOK,
			get_me_serializer.Serializer(user),
		)
	}
}
