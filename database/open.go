package database

import (
	"fmt"
	"log"
	"os"

	fmt_util "gin-rest-api/utils/fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Open() *gorm.DB {
	log.Println("Open database")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"),
		os.Getenv("MYSQL_CHARACTER_SET"),
	)

	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		log.Printf("gorm.Open() error\n%s", fmt_util.SprintfError(err))
		panic(err)
	}

	return db
}
