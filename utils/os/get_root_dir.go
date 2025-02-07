package os

import (
    "os"
    "path/filepath"
    "strings"
)

func GetRootDir() string {
    wd, _ := os.Getwd()
    for !strings.HasSuffix(wd, "/gin-rest-api") {
        wd = filepath.Dir(wd)
    }

    return wd
}
