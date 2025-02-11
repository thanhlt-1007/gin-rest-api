package token_repository

import (
	"log"
	"time"

	"gin-rest-api/initializers"
	"gin-rest-api/models"
	fmt_util "gin-rest-api/utils/fmt"
	"gin-rest-api/utils/token"
)

func CreateUserToken(user *models.User) (*models.Token, error) {
	token := models.Token{
		TokenableType: "User",
		TokenableID:   user.ID,
		AccessToken:   token.UniqueToken("access_token"),
		RefreshToken:  token.UniqueToken("refresh_token"),
		ExpiresAt:     time.Now().Add(2 * time.Hour),
	}
	tx := initializers.DB.Create(&token)
	err := tx.Error
	if err != nil {
		log.Printf("user_repository.CreateUser error\n%s", fmt_util.SprintfError(err))
		return nil, err
	}

	return &token, nil
}
