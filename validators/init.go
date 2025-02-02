package validators

import (
    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/validator/v10"
)

func Init() {
    if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
        validate.RegisterValidation("unique_in_model", UniqueInModelValidator)
    }
}
