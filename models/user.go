package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Email string `gorm:"uniqueIndex;size:255;"`
    Name *string `gorm:"size:255;"`
    EncryptedPassword string `gorm:"size:255;"`
    Tokens []Token `gorm:"polymorphic:Tokenable;constraint:OnDelete:CASCADE;"`
}
