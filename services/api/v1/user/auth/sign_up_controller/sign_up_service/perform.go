package sign_up_service

import (
    "gin-rest-api/models"
    "gin-rest-api/repositories/token_repository"
    "gin-rest-api/repositories/user_repository"
    "gin-rest-api/requests/api/v1/user/auth/sign_up_controller/sign_up_request"
    "gin-rest-api/utils/password"
    fmt_util "gin-rest-api/utils/fmt"
    "log"
)

func Perform(request sign_up_request.Request) *models.Token {
    user := models.User {
        Email: request.Email,
        EncryptedPassword: password.Encrypt(request.Password),
    }
    _, err := user_repository.CreateUser(&user)
    if err != nil {
        log.Printf("sign_up_service.Perform error\n%s", fmt_util.SprintfError(err))
        panic(err)
    }
    token, err := token_repository.CreateUserToken(&user)
    if err != nil {
        log.Printf("sign_up_service.Perform error\n%s", fmt_util.SprintfError(err))
        panic(err)
    }

    return token
}
