package authenticate_service

import (
	"log"

	"gin-rest-api/models"
	"gin-rest-api/repositories/token_repository"
	"gin-rest-api/repositories/user_repository"
	"gin-rest-api/utils/fmt_util"
)

func Perform(access_token string) (*models.User, error) {
	token, err := token_repository.GetNotExpiredToken("User", access_token)
	if err != nil {
		log.Printf("authenticate_service.Perform(access_token: %s) GetNotExpiredToken error %s\n", access_token, fmt_util.SprintfError(err))
		return nil, err
	}

	user, err := user_repository.GetUserById(token.TokenableID)
	if err != nil {
		log.Printf("authenticate_service.Perform(access_token: %s) GetUserById error %s\n", access_token, fmt_util.SprintfError(err))
		return nil, err
	}

	return user, nil
}
