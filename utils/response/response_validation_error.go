package response

import (
    "fmt"
    "gin-rest-api/responses/errors/field_error_response"
    "gin-rest-api/serializers/errors/field_error_serializer"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "net/http"
)

func ResponseValidationError(context *gin.Context, err validator.ValidationErrors) {
    var serializers []field_error_response.Response

    for _, e := range err {
        serializer := field_error_serializer.Serializer(e)
        serializers = append(serializers, serializer)
    }

    context.JSON(
        http.StatusBadRequest,
        gin.H {
            "message": fmt.Sprintf("VALIDATION_ERROR: %s", err.Error()),
            "code": "VALIDATION_ERROR",
            "errors": serializers,
        },
    )
}
