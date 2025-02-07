package database

import (
    "ariga.io/atlas-go-sdk/atlasexec"
    "context"
    "fmt"
    "log"
    "os"
    database_util "gin-rest-api/utils/database"
    fmt_util "gin-rest-api/utils/fmt"
    os_util "gin-rest-api/utils/os"
)

// FYI
// Reference: https://atlasgo.io/guides/orms/gorm/composite-types#apply-the-migrations
func Migrate() {
    log.Println("Migrate database")

    dbName := database_util.DBName()
    log.Printf("\tDB name: %s\n", dbName)

    workingDir := workingDir()
    log.Printf("\tAtlas workingDir: %s\n", workingDir)

    client, err := atlasexec.NewClient(workingDir, "atlas")
    if err != nil {
        log.Printf("atlasexec.NewClient() error\n%s", fmt_util.SprintfError(err))
        panic(err)
    }

    env := "gorm"
    url := fmt.Sprintf("mysql://%s:%s@%s/%s",
        os.Getenv("MYSQL_USER"),
        os.Getenv("MYSQL_PASSWORD"),
        os.Getenv("MYSQL_HOST"),
        dbName,
    )

    if _, err := client.MigrateApply(context.Background(), &atlasexec.MigrateApplyParams{
        Env: env,
        URL: url,
    }); err != nil {
        log.Printf("client.MigrateApply() error\n%s", fmt_util.SprintfError(err))
        panic(err)
    }
}

func workingDir() string {
    return fmt.Sprintf("%s/.", os_util.GetRootDir())
}
