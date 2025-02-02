package db

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "os"
)

func InitDB() *gorm.DB {
    db := connectDB()
    createDatabaseIfNotExists(db)
    db = openDB()

    return db
}

func connectDB() *gorm.DB {
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/",
        os.Getenv("MYSQL_USER"),
        os.Getenv("MYSQL_PASSWORD"),
        os.Getenv("MYSQL_HOST"),
    )

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    return db
}

func createDatabaseIfNotExists(db *gorm.DB) *gorm.DB {
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
    return db
}

func openDB() *gorm.DB {
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
