package sign_up_controller

import (
    "fmt"
    "gin-rest-api/requests/api/v1/user/auth/sign_up_request"
    "gin-rest-api/serializers/api/v1/user/auth/sign_up_serializer"
    "gin-rest-api/services/user/sign_up_service"
    "gin-rest-api/utils/response"
    "github.com/gin-gonic/gin"
)

func Post() gin.HandlerFunc {
    return func(context *gin.Context) {
        var request sign_up_request.Request
        err := context.ShouldBind(&request)
        if err != nil {
            fmt.Printf("SignUp err %#v - %v\n", err, err)
            panic(err)
        }

        user := sign_up_service.Perform(request)

        response.ResponseOK(
            context,
            sign_up_serializer.Serializer(user),
        )
    }
}
