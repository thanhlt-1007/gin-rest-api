package env

import (
    "fmt"
    "github.com/joho/godotenv"
    "log"
    fmt_util "gin-rest-api/utils/fmt"
    os_util "gin-rest-api/utils/os"
)

func Init() {
    log.Println("Init env")

    filename := envFileName()
    log.Printf("\tLoad %s", filename)

    err := godotenv.Load(filename)
    if err != nil {
        log.Printf("env.Init() error\n%s", fmt_util.SprintfError(err))
    }
}

func envFileName() string {
    return fmt.Sprintf("%s/.env", os_util.GetRootDir())
}
