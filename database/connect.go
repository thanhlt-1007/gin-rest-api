package database

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "os"
    fmt_util "gin-rest-api/utils/fmt"
)

func Connect() *gorm.DB {
    log.Println("Connect database")
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/",
        os.Getenv("MYSQL_USER"),
        os.Getenv("MYSQL_PASSWORD"),
        os.Getenv("MYSQL_HOST"),
    )

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Printf("gorm.Open() error\n%s", fmt_util.SprintfError(err))
        panic(err)
    }

    return db
}
