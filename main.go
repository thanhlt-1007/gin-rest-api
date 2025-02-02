package main

import (
    "gin-rest-api/database"
    "gin-rest-api/initializers"
    "gin-rest-api/middlewares"
    "gin-rest-api/router"
    "gin-rest-api/validators"
)

func init() {
    initializers.Init()
    database.Migrate()
    validators.Init()
    initializers.ENGINE.Use(middlewares.Recover())
    router.Init()
}

func main() {
    initializers.ENGINE.Run()
}
