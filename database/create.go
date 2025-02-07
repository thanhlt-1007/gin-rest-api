package database

import (
    "fmt"
    "gorm.io/gorm"
    "log"
    "os"
    database_util "gin-rest-api/utils/database"
    fmt_util "gin-rest-api/utils/fmt"
)

func Create(db *gorm.DB) {
    log.Println("Create database")

    dbName := database_util.DBName()
    log.Printf("\tDB name: %s\n", dbName)

    sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET %s COLLATE %s;",
        dbName,
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
