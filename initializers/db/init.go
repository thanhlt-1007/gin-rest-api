package db

import (
	"gin-rest-api/database"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db := database.Connect()
	database.Create(db)
	db = database.Open()

	return db
}
