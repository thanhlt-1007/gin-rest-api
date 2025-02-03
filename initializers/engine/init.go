package engine

import (
    "github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
    engine := gin.Default()

    return engine
}
