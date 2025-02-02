package database

import (
    "ariga.io/atlas-go-sdk/atlasexec"
    "context"
    "fmt"
    "log"
    "os"
)

func Migrate() {
    // FYI
    // Reference: https://atlasgo.io/guides/orms/gorm/composite-types#apply-the-migrations
    client, err := atlasexec.NewClient(".", "atlas")
    if err != nil {
        log.Fatalf("atlasexec.NewClient() error: %v", err)
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
        log.Fatalf("client.MigrateApply() error: %v", err)
    }
}
