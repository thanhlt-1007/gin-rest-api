package refresh_token_serializer

import (
	"gin-rest-api/models"
	"gin-rest-api/responses/api/v1/user/auth/refresh_token_controller/refresh_token_response"
)

func Serializer(token *models.Token) refresh_token_response.Response {
	data := refresh_token_response.Data{
		ID:           token.ID,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresAt:    token.ExpiresAt,
	}

	return refresh_token_response.Response{
		Data: data,
	}
}
