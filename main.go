package main

import (
    "gin-rest-api/initializers"
    "gin-rest-api/routers"
    "github.com/gin-gonic/gin"
)

func init() {
    initializers.InitEnv()
    initializers.InitDB()
}

func main() {
    engine := gin.Default()
    routers.SetPingRouters(engine)
    engine.Run()
}
