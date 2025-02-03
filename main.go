package main

import (
    "gin-rest-api/database"
    "gin-rest-api/initializers"
    "gin-rest-api/middlewares"
    "gin-rest-api/router"
    "gin-rest-api/validators"
    "gin-rest-api/swagger"
    "log"
)

func init() {
    initializers.Init()
    database.Migrate()
    validators.Init()
    middlewares.Init()
    swagger.Init()
    router.Init()
}

func main() {
    log.Printf("Run engine")
    initializers.ENGINE.Run()
}
