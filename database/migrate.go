package database

import (
    "ariga.io/atlas-go-sdk/atlasexec"
    "context"
    "fmt"
    "log"
    "os"
    fmt_util "gin-rest-api/utils/fmt"
)

func Migrate() {
    log.Println("Migrate database")
    // FYI
    // Reference: https://atlasgo.io/guides/orms/gorm/composite-types#apply-the-migrations
    client, err := atlasexec.NewClient(".", "atlas")
    if err != nil {
        log.Printf("atlasexec.NewClient() error\n%s", fmt_util.SprintfError(err))
        panic(err)
    }

    env := "gorm"
    url := fmt.Sprintf("mysql://%s:%s@%s/%s",
        os.Getenv("MYSQL_USER"),
        os.Getenv("MYSQL_PASSWORD"),
        os.Getenv("MYSQL_HOST"),
        os.Getenv("MYSQL_DATABASE"),
    )

    if _, err := client.MigrateApply(context.Background(), &atlasexec.MigrateApplyParams{
        Env: env,
        URL: url,
    }); err != nil {
        log.Printf("client.MigrateApply() error\n%s", fmt_util.SprintfError(err))
        panic(err)
    }
}
