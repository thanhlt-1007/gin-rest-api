package middlewares

import (
	"gin-rest-api/initializers"
)

func Init() {
	initializers.ENGINE.Use(Recover())
}
