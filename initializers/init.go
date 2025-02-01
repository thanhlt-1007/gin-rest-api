package initializers

import (
    "gin-rest-api/initializers/db"
    "gin-rest-api/initializers/engine"
    "gin-rest-api/initializers/env"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

var DB *gorm.DB
var ENGINE *gin.Engine

func Init() {
    env.InitEnv()
    DB = db.InitDB()
    ENGINE = engine.InitEngine()
}
