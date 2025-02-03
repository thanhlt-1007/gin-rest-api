package engine

import (
    "github.com/gin-gonic/gin"
    "log"
)

func Init() *gin.Engine {
    log.Println("Init engine")
    engine := gin.Default()

    return engine
}
