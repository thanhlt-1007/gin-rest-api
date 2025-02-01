package user_service

import (
    "gin-rest-api/requests/api/v1/user/auth"
    "gin-rest-api/models"
    "gin-rest-api/repositories/user_repository"
)

func SignUp(request auth.SignUpRequest) models.User {
    user := models.User {
        Email: request.Email,
    }
    user_repository.CreateUser(&user)
    return user
}
