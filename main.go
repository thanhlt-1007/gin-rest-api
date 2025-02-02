package main

import (
    "gin-rest-api/initializers"
    "gin-rest-api/middlewares"
    "gin-rest-api/migrates"
    "gin-rest-api/router"
    "gin-rest-api/validators"
)

func init() {
    initializers.Init()
    migrates.Init()
    validators.Init()
    initializers.ENGINE.Use(middlewares.Recover())
    router.Init()
}

func main() {
    initializers.ENGINE.Run()
}
