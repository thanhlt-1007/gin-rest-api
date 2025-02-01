package engine

import (
    "github.com/gin-gonic/gin"
)

func InitEngine() *gin.Engine {
    engine := gin.Default()

    return engine
}
