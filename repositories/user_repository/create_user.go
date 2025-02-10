package user_repository

import (
    "gin-rest-api/initializers"
    "gin-rest-api/models"
    "log"
    fmt_util "gin-rest-api/utils/fmt"
)

func CreateUser(user *models.User) (*models.User, error) {
    tx := initializers.DB.Create(&user)
    err := tx.Error
    if err != nil {
        log.Printf("user_repository.CreateUser error\n%s", fmt_util.SprintfError(err))
        return nil, err
    }

    return user, nil
}
