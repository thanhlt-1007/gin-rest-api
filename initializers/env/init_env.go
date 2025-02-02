package env

import (
    "github.com/joho/godotenv"
)

func InitEnv() {
    godotenv.Load()
}
