package main

import (
    "gin-rest-api/initializers"
)

func init() {
    initializers.Init()
}

func main() {
    initializers.ENGINE.Run()
}
