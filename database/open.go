package database

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "os"
)

func Open() *gorm.DB {
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
        fmt.Printf("Error [%v]\n", err)
        panic(err)
    }

    return db
}
