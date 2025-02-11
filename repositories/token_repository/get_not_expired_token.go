package token_repository

import (
	"log"
	"time"

	"gin-rest-api/initializers"
	"gin-rest-api/models"
	"gin-rest-api/utils/fmt_util"
)

func GetNotExpiredToken(tokenable_type string, access_token string) (*models.Token, error) {
	token := models.Token{}
	tx := initializers.DB.
		Where("tokenable_type = ?", tokenable_type).
		Where("access_token = ?", access_token).
		Where("expires_at > ?", time.Now()).
		First(&token)

	err := tx.Error
	if err != nil {
		log.Printf("token_repository.GetNotExpiredToken(tokenable_type: %s, access_token: %s) error\n%s", tokenable_type, access_token, fmt_util.SprintfError(err))
		return nil, err
	}

	return &token, nil
}
