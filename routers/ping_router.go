package routers

import (
    "gin-rest-api/controllers/ping_controller"
    "github.com/gin-gonic/gin"
)

func SetPingRouters(engine *gin.Engine) {
    engine.GET("/ping", ping_controller.GetPing())
}
