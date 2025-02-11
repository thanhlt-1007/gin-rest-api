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

## API doc

### Install swag

```sh
go install github.com/swaggo/swag/cmd/swag@latest
```

- Reference: https://github.com/swaggo/swag?tab=readme-ov-file#getting-started

### Init / update API doc

```sh
swag init
```

### Swagger UI

```sh
http://localhost:8080/swagger/index.html
```

## go run

```sh
go run .
```

## Format tool

### Install gofumpt

```sh
go install mvdan.cc/gofumpt@latest
```

- Reference: https://github.com/mvdan/gofumpt?tab=readme-ov-file#gofumpt

### gofumpt

```sh
gofumpt -e -l -w .
```

### Lint tool

### Install golangci-lint

```sh
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63.4
```

### golangci-lint

```sh
golangci-lint run
```
