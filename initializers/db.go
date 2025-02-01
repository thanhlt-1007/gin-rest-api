package initializers

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "os"
)

var DB *gorm.DB

func InitDB() {
    fmt.Println("\n---INIT-DB---")
    db, err := gorm.Open(
        sqlite.Open(os.Getenv("DB_NAME")),
        &gorm.Config{},
    )
    if err != nil {
        fmt.Printf("Init DB error [%v]\n", err)
        panic(err)
    }

    DB = db

    fmt.Println("Init db success")
}
