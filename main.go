package main

import (
    "gin-rest-api/initializers"
    "gin-rest-api/router"
    "gin-rest-api/middlewares"
)

func init() {
    initializers.Init()
    initializers.ENGINE.Use(middlewares.Recover())
    router.Init()
}

func main() {
    initializers.ENGINE.Run()
}
