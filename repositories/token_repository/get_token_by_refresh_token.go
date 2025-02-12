package token_repository

import (
	"log"

	"gin-rest-api/initializers"
	"gin-rest-api/models"
	"gin-rest-api/utils/fmt_util"
)

func GetTokenByRefreshToken(refresh_token string) (*models.Token, error) {
	var token models.Token
	tx := initializers.DB.
		Where("refresh_token = ?", refresh_token).
		First(&token)
	err := tx.Error
	if err != nil {
		log.Printf("token_repository.GetTokenByRefreshToken(refresh_token: %s) error\n%s", refresh_token, fmt_util.SprintfError(err))
		return nil, err
	}

	return &token, nil
}
