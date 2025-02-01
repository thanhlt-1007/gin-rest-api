package initializers

import (
    "github.com/joho/godotenv"
    "fmt"
)

func InitEnv() {
    fmt.Println("\n---INIT-ENV---")

    err := godotenv.Load()
    if err != nil {
        fmt.Printf("Init env error [%v], skip load ENV from .env\n", err)
    }

    fmt.Println("Init env success")
}
