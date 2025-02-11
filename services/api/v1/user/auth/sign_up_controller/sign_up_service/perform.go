package sign_up_service

import (
	"log"

	"gin-rest-api/models"
	"gin-rest-api/repositories/token_repository"
	"gin-rest-api/repositories/user_repository"
	"gin-rest-api/requests/api/v1/user/auth/sign_up_controller/sign_up_request"
	"gin-rest-api/utils/fmt_util"
	"gin-rest-api/utils/password_util"
)

func Perform(request sign_up_request.Request) *models.Token {
	user := models.User{
		Email:             request.Email,
		EncryptedPassword: password_util.Encrypt(request.Password),
	}
	_, createUserErr := user_repository.CreateUser(&user)
	if createUserErr != nil {
		log.Printf("sign_up_service.Perform CreateUser error\n%s", fmt_util.SprintfError(createUserErr))
		panic(createUserErr)
	}
	token, createUserTokenErr := token_repository.CreateUserToken(&user)
	if createUserTokenErr != nil {
		log.Printf("sign_up_service.Perform CreateUserToken error\n%s", fmt_util.SprintfError(createUserTokenErr))
		panic(createUserTokenErr)
	}

	return token
}
