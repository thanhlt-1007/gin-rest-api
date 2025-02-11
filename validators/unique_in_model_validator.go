package validators

import (
	"errors"
	"fmt"
	"reflect"

	"gin-rest-api/initializers"
	"gin-rest-api/models"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

var UniqueInModelValidator validator.Func = func(fieldLevel validator.FieldLevel) bool {
	fieldName := strcase.ToSnake(fieldLevel.FieldName())
	fieldValue := fieldLevel.Field()
	model := fieldLevel.Param()

	modelInterface := modelToInterface(model)
	queryValue := fieldValueToQueryValue(fieldValue)

	var count int64
	query := fmt.Sprintf("%s = ?", fieldName)
	initializers.DB.Model(modelInterface).Where(query, queryValue).Count(&count)

	return count <= 0
}

func modelToInterface(model string) any {
	switch model {
	case "User":
		return &models.User{}
	default:
		message := fmt.Sprintf("UniqueInModelValidator modelToInterface not implemented for model: %s", model)
		panic(errors.New(message))
	}
}

func fieldValueToQueryValue(fieldValue reflect.Value) any {
	switch fieldValue.Kind() {
	case reflect.String:
		return fieldValue.String()
	default:
		message := fmt.Sprintf("UniqueInModelValidator fieldValueToQueryValue not implemented for fieldValue: %v", fieldValue)
		panic(errors.New(message))
	}
}
