package token_repository

import (
    "gin-rest-api/initializers"
    "gin-rest-api/models"
    "gin-rest-api/utils/token"
    "time"
)

func CreateUserToken(user models.User) models.Token {
    token := models.Token {
        TokenableType: "User",
        TokenableID: user.ID,
        AccessToken: token.UniqueToken("access_token"),
        RefreshToken: token.UniqueToken("refresh_token"),
        ExpiresAt: time.Now().Add(2 * time.Hour),
    }
    result := initializers.DB.Create(&token)

    err := result.Error
    if err != nil {
        panic(err)
    }

    return token
}
