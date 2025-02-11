package engine

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	log.Println("Init engine")
	engine := gin.Default()

	return engine
}
