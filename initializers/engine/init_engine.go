package engine

import (
    "gin-rest-api/routers"
    "github.com/gin-gonic/gin"
)

func InitEngine() *gin.Engine {
    engine := gin.Default()

    // API /ping
    routers.SetPingRouters(engine)

    return engine
}
