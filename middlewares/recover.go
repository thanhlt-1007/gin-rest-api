package middlewares

import (
    "fmt"
    "gin-rest-api/utils/response"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "net/http"
)

func Recover() gin.HandlerFunc {
    return func(context *gin.Context) {
        defer func ()  {
            recovered := recover()
            if recovered != nil {
                if err, ok := recovered.(validator.ValidationErrors); ok {
                    response.ResponseValidationError(context, err)
                } else if err, ok := recovered.(error); ok {
                    context.JSON(
                        http.StatusInternalServerError,
                        gin.H {
                            "message": fmt.Sprintf("UNKNOW_ERROR: %s", err.Error()),
                            "code": "INTERNAL_SERVER_ERROR",
                        },
                    )
                } else {
                    context.JSON(
                        http.StatusInternalServerError,
                        gin.H {
                            "message": fmt.Sprintf("UNKNOW_PANIC: %#v - %v", recovered, recovered),
                            "code": "INTERNAL_SERVER_ERROR",
                        },
                    )
                }
                context.Abort()
            }
        }()

        context.Next()
    }
}
