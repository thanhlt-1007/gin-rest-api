package user_service

import (
    "gin-rest-api/models"
    "gin-rest-api/repositories/user_repository"
    "gin-rest-api/requests/api/v1/user/auth"
    "gin-rest-api/utils/password"
)

func SignUp(request auth.SignUpRequest) models.User {
    user := models.User {
        Email: request.Email,
        EncryptedPassword: password.Encrypt(request.Password),
    }
    user_repository.CreateUser(&user)
    return user
}
