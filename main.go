package main

import (
    "gin-rest-api/database"
    "gin-rest-api/initializers"
    "gin-rest-api/middlewares"
    "gin-rest-api/router"
    "gin-rest-api/validators"
    "log"
)

func init() {
    initializers.Init()
    database.Migrate()
    validators.Init()
    middlewares.Init()
    router.Init()
}

func main() {
    log.Printf("Run engine")
    initializers.ENGINE.Run()
}
