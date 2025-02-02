package db

import (
    "fmt"
    "gin-rest-api/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "os"
)

func InitDB() *gorm.DB {
    fmt.Println("\n---INIT-DB---")

    db := openDB()
    migrateDB(db)

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

func migrateDB(db *gorm.DB) {
    fmt.Println("\n---MIGRATE-DB---")

    err := db.AutoMigrate(&models.User{}, &models.Token{})
    if err != nil {
        fmt.Printf("Error [%v]\n", err)
        panic(err)
    }

    fmt.Println("\n---MIGRATE-DB-SUCCESS---")
}
