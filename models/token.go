package models

import (
    "time"
)

type Token struct {
    ID uint `gorm:"primaryKey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    TokenableID uint
    TokenableType string

    AccessToken string `gorm:"uniqueIndex"`
    RefreshToken string `gorm:"uniqueIndex"`
    ExpiresAt time.Time
}
