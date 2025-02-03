package db

import (
    "gorm.io/gorm"
    "gin-rest-api/database"
)

func Init() *gorm.DB {
    db := database.Connect()
    database.Create(db)
    db = database.Open()

    return db
}
