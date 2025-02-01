package auth

import (
    "gin-rest-api/models"
)

type SignUpResponse struct {
    ID uint `json:"id"`
    Email string `json:"email"`
}

func SignUpSerializer(user models.User) SignUpResponse {
    return SignUpResponse {
        ID: user.ID,
        Email: user.Email,
    }
}
