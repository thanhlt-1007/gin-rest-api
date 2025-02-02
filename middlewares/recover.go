package middlewares

import (
    "fmt"
    "gin-rest-api/responses/errors/field_error_response"
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

                    var serializers []field_error_response.Response

                    for _, fieldError := range err {
                        serializer := field_error_response.Response {
                            Field: fieldError.Field(),
                            Tag: fieldError.Tag(),
                        }
                        serializers = append(serializers, serializer)
                    }

                    context.JSON(
                        http.StatusBadRequest,
                        gin.H {
                            "message": fmt.Sprintf("VALIDATION_ERROR: %s", err.Error()),
                            "code": "VALIDATION_ERROR",
                            "errors": serializers,
                        },
                    )
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
