package sign_in_response

import (
	"time"
)

type Response struct {
	Data Data `json:"data"`
}

type Data struct {
	ID           uint      `json:"id" example:"1"`
	AccessToken  string    `json:"access_token" example:"xxx"`
	RefreshToken string    `json:"refresh_token" example:"xxx"`
	ExpiresAt    time.Time `json:"expires_at" example:"2025-12-31T12:30:00.0000+07:00"`
}
