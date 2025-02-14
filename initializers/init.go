package initializers

import (
	"gin-rest-api/initializers/db"
	"gin-rest-api/initializers/engine"
	"gin-rest-api/initializers/env"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	ENGINE *gin.Engine
)

func Init() {
	env.Init()
	DB = db.Init()
	ENGINE = engine.Init()
}
