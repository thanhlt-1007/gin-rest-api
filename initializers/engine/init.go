package engine

import (
    "github.com/gin-gonic/gin"
    "log"
)

func Init() *gin.Engine {
    log.Println("Init engine")

    gin.SetMode(gin.ReleaseMode)
    engine := gin.Default()

    return engine
}
