# gorm-rest-api

## gvm

```sh
gvm use go1.23.5
```

## go get

```sh
go get .
```

## atlas

### Install

```sh
curl -sSf https://atlasgo.sh | sh
```

- Reference: https://atlasgo.io/guides/orms/gorm/standalone#installation

### Generate Migrations For the Schema

```sh
atlas migrate diff --env gorm --dev-url 'mysql://{MYSQL_USER}:{MYSQL_PASSWORD}@{MYSQL_HOST}/{MYSQL_DATABASE}'
```

Example:

```sh
atlas migrate diff --env gorm --dev-url 'mysql://root:Aa@123456@localhost:3306/gin_rest_api'
```

- Reference: https://atlasgo.io/guides/orms/gorm/composite-types#generate-migrations-for-the-schema

## go run

```sh
go run .
```
