package middlewares

import (
	"gin-rest-api/utils/errors_util"
	"gin-rest-api/utils/response/response_internal_server_error"
	"gin-rest-api/utils/response/response_unique_contraint_error"
	"gin-rest-api/utils/response/response_unknown_panic_error"
	"gin-rest-api/utils/response/response_validation_error"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Recovery() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			recovered := recover()
			if recovered != nil {
				if err, ok := recovered.(validator.ValidationErrors); ok {
					response_validation_error.JSON(context, err)
				} else if err, ok := recovered.(error); ok {
					// FYI
					// Hanlde `gorm:"uniqueIndex"`
					if uniqueContraintError, ok := errors_util.IsUniqueContraintError(err); ok {
						response_unique_contraint_error.JSON(context, *uniqueContraintError)
					} else {
						response_internal_server_error.JSON(context, err)
					}
				} else {
					response_unknown_panic_error.JSON(context, recovered)
				}
				context.Abort()
			}
		}()

		context.Next()
	}
}
