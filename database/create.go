package database

import (
	"fmt"
	"log"
	"os"

	"gin-rest-api/utils/fmt_util"

	"gorm.io/gorm"
)

func Create(db *gorm.DB) {
	log.Println("Create database")
	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET %s COLLATE %s;",
		os.Getenv("MYSQL_DATABASE"),
		os.Getenv("MYSQL_CHARACTER_SET"),
		os.Getenv("MYSQL_COLLATE"),
	)
	db = db.Exec(sql)
	err := db.Error
	if err != nil {
		log.Printf("db.Exec() error\n%s", fmt_util.SprintfError(err))
		panic(err)
	}
}
