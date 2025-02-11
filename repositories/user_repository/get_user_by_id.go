package user_repository

import (
	"log"

	"gin-rest-api/initializers"
	"gin-rest-api/models"
	"gin-rest-api/utils/fmt_util"
)

func GetUserById(id uint) (*models.User, error) {
	user := models.User{}

	tx := initializers.DB.Where("id = ?", id).First(&user)
	err := tx.Error
	if err != nil {
		log.Printf("user_repository.GetUserById(%d) error\n%s", id, fmt_util.SprintfError(err))
		return nil, err
	}

	return &user, nil
}
