package db

import (
    "fmt"
    "gin-rest-api/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "os"
)

var DB *gorm.DB

func InitDB() {
    fmt.Println("\n---INIT-DB---")

    openDB()
    migrateDB()

    fmt.Println("\n---INIT-DB-SUCCESS---")
}

func openDB() {
    fmt.Println("\n---OPEN-DB---")

    db, err := gorm.Open(
        sqlite.Open(os.Getenv("DB_NAME")),
        &gorm.Config{},
    )
    if err != nil {
        fmt.Printf("Error [%v]\n", err)
        panic(err)
    }

    DB = db

    fmt.Println("\n---OPEN-DB-SUCCESS---")
}

func migrateDB() {
    fmt.Println("\n---MIGRATE-DB---")

    err := DB.AutoMigrate(&models.User{})
    if err != nil {
        fmt.Printf("Error [%v]\n", err)
        panic(err)
    }

    fmt.Println("\n---MIGRATE-DB-SUCCESS---")
}
