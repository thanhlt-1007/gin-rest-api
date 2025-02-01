package initializers

import (
    "gin-rest-api/initializers/db"
    "gin-rest-api/initializers/env"
)

func Init() {
    env.InitEnv()
    db.InitDB()
}
