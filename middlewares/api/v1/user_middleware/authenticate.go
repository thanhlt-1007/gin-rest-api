package user_middleware

import (
	"gin-rest-api/errors"
	"gin-rest-api/services/middlewares/api/v1/user_middleware/authenticate_service"

	"gin-rest-api/consts/middlewares/api/v1/user_middleware/authenticate_consts"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(context *gin.Context) {
		access_token := context.GetHeader("X-Access-Token")
		user, err := authenticate_service.Perform(access_token)
		if err != nil {
			context.Abort()
			panic(errors.UnauthorizedError{
				RootError: err,
			})
		}

		context.Set(authenticate_consts.AUTHENTICATED_USER, user)
		context.Next()
	}
}
