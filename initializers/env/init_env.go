package env

import (
    "github.com/joho/godotenv"
    "fmt"
)

func InitEnv() {
    fmt.Println("\n---INIT-ENV---")

    err := godotenv.Load()
    if err != nil {
        fmt.Printf("Error [%v], skip load ENV from .env\n", err)
    }

    fmt.Println("\n---INIT-ENV-SUCCESS---")
}
