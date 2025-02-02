package sign_up_service

import (
    "gin-rest-api/models"
    "gin-rest-api/repositories/user_repository"
    "gin-rest-api/requests/api/v1/user/auth/sign_up_request"
    "gin-rest-api/utils/password"
)

func Perform(request sign_up_request.Request) models.User {
    user := models.User {
        Email: request.Email,
        EncryptedPassword: password.Encrypt(request.Password),
    }
    user_repository.CreateUser(&user)
    return user
}
