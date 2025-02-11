package sign_in_serializer

import (
	"gin-rest-api/models"
	"gin-rest-api/responses/api/v1/user/auth/sign_in_controller/sign_in_response"
)

func Serializer(token *models.Token) sign_in_response.Response {
	data := sign_in_response.Data{
		ID:           token.ID,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresAt:    token.ExpiresAt,
	}

	return sign_in_response.Response{
		Data: data,
	}
}
