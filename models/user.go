package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Email string `gorm:"uniqueIndex"`
    Name *string
    EncryptedPassword string
    Tokens []Token `gorm:"polymorphic:Tokenable;constraint:OnDelete:CASCADE;"`
}
