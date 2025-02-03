package validators

import (
    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/validator/v10"
    "log"
)

func Init() {
    log.Println("Init validators")
    if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
        validate.RegisterValidation("unique_in_model", UniqueInModelValidator)
    }
}
