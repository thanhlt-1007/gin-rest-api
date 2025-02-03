package swagger

import (
    "gin-rest-api/initializers"
    "log"
    docs "gin-rest-api/docs"
    swaggo_files "github.com/swaggo/files"
    swaggo_gin_swagger "github.com/swaggo/gin-swagger"
)

func Init() {
    log.Println("Init swagger")
    docs.SwaggerInfo.Title = "Gin Rest API Swagger UI"
    docs.SwaggerInfo.Version = "1.0"
    docs.SwaggerInfo.Description = "This is API doc for Gin Rest API."
    docs.SwaggerInfo.Schemes = []string {"http", "https"}
    docs.SwaggerInfo.Host = "localhost:8080"
    docs.SwaggerInfo.BasePath = "/api/v1"

    initializers.ENGINE.GET("/swagger/*any", swaggo_gin_swagger.WrapHandler(swaggo_files.Handler))
}
