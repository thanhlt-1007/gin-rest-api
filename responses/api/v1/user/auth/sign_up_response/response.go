package sign_up_response

import (
    "time"
)

type Response struct {
    ID uint `json:"id"`
    AccessToken string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
    ExpiresAt time.Time `json:"expires_at"`
}
