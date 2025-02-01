package main

import (
    "gin-rest-api/initializers"
    "gin-rest-api/router"
)

func init() {
    initializers.Init()
    router.Init()
}

func main() {
    initializers.ENGINE.Run()
}
