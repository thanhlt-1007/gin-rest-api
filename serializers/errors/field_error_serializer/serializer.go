package field_error_serializer

import (
	"gin-rest-api/responses/errors/field_error_response"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

func Serializer(err validator.FieldError) field_error_response.Response {
	field := strcase.ToSnake(err.Field())
	tag := err.Tag()

	return field_error_response.Response{
		Field: field,
		Tag:   tag,
	}
}
