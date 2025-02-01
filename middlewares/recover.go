package middlewares

import (
    "fmt"
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
                    fmt.Printf("err: %v\n", err)
                    fmt.Printf("err: %#v\n", err)

                    context.JSON(
                        http.StatusBadRequest,
                        gin.H {
                            "error": fmt.Sprintf("VALIDATION_ERRORS: %s", err.Error()),
                        },
                    )
                } else if err, ok := recovered.(error); ok {
                    context.JSON(
                        http.StatusInternalServerError,
                        gin.H {
                            "error": fmt.Sprintf("UNKNOW_ERRORS: %s", err.Error()),
                        },
                    )
                } else {
                    context.JSON(
                        http.StatusInternalServerError,
                        gin.H {
                            "error": fmt.Sprintf("UNKNOW_PANIC: %#v - %v", recovered, recovered),
                        },
                    )
                }
                context.Abort()
            }
        }()

        context.Next()
    }
}
