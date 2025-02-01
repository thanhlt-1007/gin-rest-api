package main

import (
    "gin-rest-api/routers"
    "github.com/gin-gonic/gin"
)

func main() {
    engine := gin.Default()

    routers.SetPingRouters(engine)

    engine.Run()
}
