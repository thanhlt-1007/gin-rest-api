package middlewares

import (
    "gin-rest-api/initializers"
    "log"
)

func Init() {
    log.Println("Init middlewares")
    initializers.ENGINE.Use(Recover())
}
