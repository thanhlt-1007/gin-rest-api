package db

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "os"
)

func InitDB() *gorm.DB {
    fmt.Println("\n---INIT-DB---")

    db := openDB()

    fmt.Println("\n---INIT-DB-SUCCESS---")
    return db
}

func openDB() *gorm.DB {
    fmt.Println("\n---OPEN-DB---")

    db, err := gorm.Open(
        sqlite.Open(os.Getenv("DB_NAME")),
        &gorm.Config{
            Logger: logger.Default.LogMode(logger.Info),
        },
    )
    if err != nil {
        fmt.Printf("Error [%v]\n", err)
        panic(err)
    }

    fmt.Println("\n---OPEN-DB-SUCCESS---")
    return db
}
