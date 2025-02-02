package migrates

import (
    "fmt"
    "gin-rest-api/initializers"
    "gin-rest-api/models"
)

func Init() {
    err := initializers.DB.AutoMigrate(&models.User{}, &models.Token{})
    if err != nil {
        fmt.Printf("Error [%v]\n", err)
        panic(err)
    }
}
