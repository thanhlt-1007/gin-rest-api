package middlewares

import (
	"log"

	"gin-rest-api/utils/errors_util"
	"gin-rest-api/utils/response_util"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Recovery() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			recovered := recover()
			if recovered != nil {
				log.Printf("Recovery recovered %#v\n", recovered)

				if err, ok := recovered.(validator.ValidationErrors); ok {
					response_util.ResponseValidationError(context, err)
				} else if err, ok := recovered.(error); ok {
					// FYI
					// Hanlde `gorm:"uniqueIndex"`
					if uniqueContraintError, ok := errors_util.IsUniqueContraintError(err); ok {
						response_util.ResponseUniqueContraintError(context, *uniqueContraintError)
					} else {
						response_util.ResponseInternalServerError(context, err)
					}
				} else {
					response_util.ResponseUnknownPanicError(context, recovered)
				}
				context.Abort()
			}
		}()

		context.Next()
	}
}
