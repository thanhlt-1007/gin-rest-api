package middlewares

import (
    "gin-rest-api/utils/response"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
)

func Recover() gin.HandlerFunc {
    return func(context *gin.Context) {
        defer func ()  {
            recovered := recover()
            if recovered != nil {
                if err, ok := recovered.(validator.ValidationErrors); ok {
                    response.ResponseValidationError(context, err)
                } else if err, ok := recovered.(error); ok {
                    response.ResponseInternalServerError(context, err)
                } else {
                    response.ResponseUnkownPanicError(context, recovered)
                }
                context.Abort()
            }
        }()

        context.Next()
    }
}
