package refresh_token_service

import (
	"log"

	"gin-rest-api/errors"
	"gin-rest-api/models"
	"gin-rest-api/repositories/token_repository"
	"gin-rest-api/repositories/user_repository"
	"gin-rest-api/utils/fmt_util"
)

func Perform(refresh_token string) *models.Token {
	token, err := token_repository.GetTokenByRefreshToken(refresh_token)
	if err != nil {
		log.Printf("refresh_token_service.Perform(refresh_token: %s) GetTokenByRefreshToken error\n%s", refresh_token, fmt_util.SprintfError(err))
		panic(errors.UnauthorizedError{
			RootError: err,
		})
	}

	user, err := user_repository.GetUserById(token.TokenableID)
	if err != nil {
		log.Printf("refresh_token_service.Perform(refresh_token: %s) GetUserById error\n%s", refresh_token, fmt_util.SprintfError(err))
		panic(errors.UnauthorizedError{
			RootError: err,
		})
	}

	_, err = token_repository.DeleteToken(token)
	if err != nil {
		log.Printf("refresh_token_service.Perform(refresh_token: %s) DeleteToken error\n%s", refresh_token, fmt_util.SprintfError(err))
		panic(errors.UnauthorizedError{
			RootError: err,
		})
	}

	newToken, err := token_repository.CreateUserToken(user)
	if err != nil {
		log.Printf("refresh_token_service.Perform(refresh_token: %s) CreateUserToken error\n%s", refresh_token, fmt_util.SprintfError(err))
		panic(errors.UnauthorizedError{
			RootError: err,
		})
	}

	return newToken
}
