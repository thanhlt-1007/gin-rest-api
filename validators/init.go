package validators

import (
	"log"

	fmt_util "gin-rest-api/utils/fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Init() {
	log.Println("Init validators")
	if validate, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := validate.RegisterValidation("unique_in_model", UniqueInModelValidator)
		if err != nil {
			log.Printf("validate.RegisterValidation('unique_in_model') error\n%s", fmt_util.SprintfError(err))
			panic(err)
		}
	}
}
