package field_error_serializer

import (
    "gin-rest-api/responses/errors/field_error_response"
    "github.com/go-playground/validator/v10"
    "github.com/iancoleman/strcase"
)

func Serializer(err validator.FieldError) field_error_response.Response {
    return field_error_response.Response {
        Field: strcase.ToSnake(err.Field()),
        Tag: err.Tag(),
    }
}
