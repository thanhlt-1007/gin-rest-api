package env

import (
	"log"

	"gin-rest-api/utils/fmt_util"

	"github.com/joho/godotenv"
)

func Init() {
	log.Println("Init env")
	err := godotenv.Load()
	if err != nil {
		log.Printf("env.Init() error\n%s", fmt_util.SprintfError(err))
	}
}
