package get_me_serializer

import (
	"gin-rest-api/models"
	"gin-rest-api/responses/api/v1/user/me_controller/get_me_response"
)

func Serializer(user *models.User) get_me_response.Response {
	data := get_me_response.Data{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	return get_me_response.Response{
		Data: data,
	}
}
