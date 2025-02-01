package auth_controller

import (
    "fmt"
    "gin-rest-api/services/user_service"
    "github.com/gin-gonic/gin"
    "net/http"
    request "gin-rest-api/requests/api/v1/user/auth"
    serializer "gin-rest-api/serializers/api/v1/user/auth"
)

func SignUp() gin.HandlerFunc {
    return func(context *gin.Context) {
        var signUpRequest request.SignUpRequest
        err := context.ShouldBind(&signUpRequest)
        if err != nil {
            fmt.Printf("SignUp err %#v - %v\n", err, err)
            context.JSON(
                http.StatusUnprocessableEntity,
                gin.H {
                    "error": err.Error(),
                },
            )
            return
        }

        user := user_service.SignUp(signUpRequest)

        context.JSON(
            http.StatusOK,
            serializer.Serializer(user),
        )
    }
}
