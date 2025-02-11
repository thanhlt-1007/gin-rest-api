package env

import (
	"log"

	fmt_util "gin-rest-api/utils/fmt"
	"github.com/joho/godotenv"
)

func Init() {
	log.Println("Init env")
	err := godotenv.Load()
	if err != nil {
		log.Printf("env.Init() error\n%s", fmt_util.SprintfError(err))
	}
}
