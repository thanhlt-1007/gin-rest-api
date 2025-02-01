package user_repository

import (
    "gin-rest-api/models"
    "gin-rest-api/initializers"
)

func CreateUser(user *models.User) {
    result := initializers.DB.Create(&user)
    err := result.Error
    if err != nil {
        panic(err)
    }
}
