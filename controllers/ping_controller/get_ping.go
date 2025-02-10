package ping_controller

import (
    _ "gin-rest-api/responses/ping_controller/get_ping_response"
    "gin-rest-api/serializers/ping_controller/get_ping_serializer"
    "gin-rest-api/utils/response/response_ok"
    "github.com/gin-gonic/gin"
)

// godoc
// @Tags Ping
// @Router /ping [get]
// @Summary GET ping used for healthcheck
// @Accept json
// @Produce json
// @Success 200	{object} get_ping_response.Response
func GetPing() gin.HandlerFunc {
    return func(context *gin.Context) {
        response_ok.JSON(
            context,
            get_ping_serializer.Serializer(),
        )
    }
}
