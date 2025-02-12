package token_repository

import (
	"gin-rest-api/initializers"
	"gin-rest-api/models"
)

func DeleteToken(token *models.Token) (bool, error) {
	tx := initializers.DB.Delete(&token)

	err := tx.Error
	if err != nil {
		return false, err
	}

	return true, err
}
