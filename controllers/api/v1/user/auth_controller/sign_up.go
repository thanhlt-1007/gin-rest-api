package auth_controller

import (
    "fmt"
    "gin-rest-api/services/user_service"
    "gin-rest-api/utils/response"
    "github.com/gin-gonic/gin"
    request "gin-rest-api/requests/api/v1/user/auth"
    serializer "gin-rest-api/serializers/api/v1/user/auth"
)

func SignUp() gin.HandlerFunc {
    return func(context *gin.Context) {
        var signUpRequest request.SignUpRequest
        err := context.ShouldBind(&signUpRequest)
        if err != nil {
            fmt.Printf("SignUp err %#v - %v\n", err, err)
            panic(err)
        }

        user := user_service.SignUp(signUpRequest)

        response.ResponseOK(
            context,
            serializer.SignUpSerializer(user),
        )
    }
}
