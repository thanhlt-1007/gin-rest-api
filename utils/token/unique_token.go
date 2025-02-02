package token

import (
    "fmt"
    "gin-rest-api/initializers"
    "gin-rest-api/models"
)

func UniqueToken(field string) string {
    token := RandomToken()
    var count int64
    query := fmt.Sprintf("%s = ?", field)
    initializers.DB.Model(&models.Token{}).Where(query, token).Count(&count)
    if count == 0 {
        return token
    }

    return UniqueToken(field)
}
