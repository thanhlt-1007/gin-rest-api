package database

import (
    "fmt"
    "os"
)

func DBName() string {
    name := os.Getenv("MYSQL_DATABASE")
    if os.Getenv("GO_TEST") == "true" {
        name = fmt.Sprintf("%s_test", name)
    }
    return name
}
