package validation_error_serializer

import (
    "fmt"
    "gin-rest-api/consts"
    "gin-rest-api/responses/errors/field_error_response"
    "gin-rest-api/responses/errors/validation_error_response"
    "gin-rest-api/serializers/errors/field_error_serializer"
    "github.com/go-playground/validator/v10"
)

func Serializer(err validator.ValidationErrors) validation_error_response.Response {
    message := fmt.Sprintf("VALIDATION_ERROR: %s", err.Error())
    code := consts.VALIDATION_ERROR_CODE

    var errors []field_error_response.Response
    for _, e := range err {
        err := field_error_serializer.Serializer(e)
        errors = append(errors, err)
    }

    return validation_error_response.Response {
        Message: message,
        Code: code,
        Errors: errors,
    }
}
