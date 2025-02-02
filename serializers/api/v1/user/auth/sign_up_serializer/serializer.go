package sign_up_serializer

import (
    "gin-rest-api/models"
    "gin-rest-api/responses/api/v1/user/auth/sign_up_response"
)

func Serializer(token models.Token) sign_up_response.Response {
    return sign_up_response.Response {
        ID: token.ID,
        AccessToken: token.AccessToken,
        RefreshToken: token.RefreshToken,
        ExpiresAt: token.ExpiresAt,
    }
}
