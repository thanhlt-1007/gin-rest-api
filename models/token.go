package models

import (
	"time"
)

type Token struct {
	ID            uint `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	TokenableID   uint
	TokenableType string `gorm:"size:255;"`
	AccessToken   string `gorm:"uniqueIndex;size:255;"`
	RefreshToken  string `gorm:"uniqueIndex;size:255;"`
	ExpiresAt     time.Time
}
