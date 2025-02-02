package middlewares

import (
    "gin-rest-api/utils/response/response_internal_server_error"
    "gin-rest-api/utils/response/response_unkown_panic_error"
    "gin-rest-api/utils/response/response_validation_error"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
)

func Recover() gin.HandlerFunc {
    return func(context *gin.Context) {
        defer func ()  {
            recovered := recover()
            if recovered != nil {
                if err, ok := recovered.(validator.ValidationErrors); ok {
                    response_validation_error.JSON(context, err)
                } else if err, ok := recovered.(error); ok {
                    response_internal_server_error.JSON(context, err)
                } else {
                    response_unkown_panic_error.JSON(context, recovered)
                }
                context.Abort()
            }
        }()

        context.Next()
    }
}
