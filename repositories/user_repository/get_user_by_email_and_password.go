package user_repository

import (
	"log"

	"gin-rest-api/initializers"
	"gin-rest-api/models"
	"gin-rest-api/utils/fmt_util"
	"gin-rest-api/utils/password_util"
)

func GetUserByEmailAndPassword(email string, password string) (*models.User, error) {
	user := models.User{}

	tx := initializers.DB.Where("email = ?", email).First(&user)
	firstUserErr := tx.Error
	if firstUserErr != nil {
		log.Printf("user_repository.GetUserByEmailAndPassword find first user by email %s error\n%s", email, fmt_util.SprintfError(firstUserErr))
		return nil, firstUserErr
	}

	_, compareErr := password_util.Compare(user.EncryptedPassword, password)
	if compareErr != nil {
		log.Printf("user_repository.GetUserByEmailAndPassword compare password error\n%s", fmt_util.SprintfError(compareErr))
		return nil, compareErr
	}

	return &user, nil
}
