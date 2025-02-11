package sign_in_service

import (
	"gin-rest-api/errors"
	"gin-rest-api/models"
	"gin-rest-api/repositories/token_repository"
	"gin-rest-api/repositories/user_repository"
	"gin-rest-api/requests/api/v1/user/auth/sign_in_controller/sign_in_request"
)

func Perform(request sign_in_request.Request) *models.Token {
	user, err := user_repository.GetUserByEmailAndPassword(request.Email, request.Password)
	if err != nil {
		panic(errors.UnauthorizedError{
			RootError: err,
		})
	}

	token, err := token_repository.CreateUserToken(user)
	if err != nil {
		panic(errors.UnauthorizedError{
			RootError: err,
		})
	}

	return token
}
