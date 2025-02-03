package ping_controller

import (
    _ "gin-rest-api/swagger/responses/ping_controller/get_ping"
    "gin-rest-api/serializers/ping_serializer"
    "gin-rest-api/utils/response/response_ok"
    "github.com/gin-gonic/gin"
)

// godoc
// @Tags Ping
// @Router /ping [get]
// @Summary GET ping used for healthcheck
// @Accept json
// @Produce json
// @Success 200	{object} get_ping.SuccessResponse
func GetPing() gin.HandlerFunc {
    return func(context *gin.Context) {
        response_ok.JSON(
            context,
            ping_serializer.Serializer(),
        )
    }
}
