package ping_controller

import (
	// "net/http"
	"gin-rest-api/initializers"
	"net/http"
	"net/http/httptest"
	"testing"

	// "github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	// "gin-rest-api/main"
	"gin-rest-api/database"
	// "gin-rest-api/initializers"
	"gin-rest-api/middlewares"
	"gin-rest-api/router"
	"gin-rest-api/swagger"
	"gin-rest-api/validators"
	// "log"
    "os"
)

func TestGetPing(test *testing.T) {
    initTest()

    writer := httptest.NewRecorder()
    request, _ := http.NewRequest("GET", "/ping", nil)
    initializers.ENGINE.ServeHTTP(writer, request)

    assert.Equal(test, 200, writer.Code)
}

func initTest() {
    os.Setenv("GO_TEST", "true")

    // gin.SetMode(gin.TestMode)
    initializers.Init()
    database.Migrate()
    validators.Init()
    middlewares.Init()
    swagger.Init()
    router.Init()
}