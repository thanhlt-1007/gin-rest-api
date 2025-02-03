package sign_up_controller

import (
    "gin-rest-api/requests/api/v1/user/auth/sign_up_request"
    "gin-rest-api/serializers/api/v1/user/auth/sign_up_serializer"
    "gin-rest-api/services/user/sign_up_service"
    "gin-rest-api/utils/response/response_ok"
    "github.com/gin-gonic/gin"
)

func PostSignUp() gin.HandlerFunc {
    return func(context *gin.Context) {
        var request sign_up_request.Request
        err := context.ShouldBind(&request)
        if err != nil {
            panic(err)
        }

        token := sign_up_service.Perform(request)

        response_ok.JSON(
            context,
            sign_up_serializer.Serializer(token),
        )
    }
}
