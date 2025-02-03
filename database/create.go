package database

import (
    "fmt"
    "gorm.io/gorm"
    "os"
)

func Create(db *gorm.DB) {
    sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET %s COLLATE %s;",
        os.Getenv("MYSQL_DATABASE"),
        os.Getenv("MYSQL_CHARACTER_SET"),
        os.Getenv("MYSQL_COLLATE"),
    )
    db = db.Exec(sql)

    err := db.Error
    if err != nil {
        panic(err)
    }
}
