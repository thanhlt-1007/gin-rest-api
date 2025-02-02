package main

import (
    "gin-rest-api/initializers"
    "gin-rest-api/middlewares"
    "gin-rest-api/router"
    "gin-rest-api/validators"
)

func init() {
    initializers.Init()
}

func main() {
    validators.Init()
    initializers.ENGINE.Use(middlewares.Recover())
    router.Init()
    initializers.ENGINE.Run()
}
