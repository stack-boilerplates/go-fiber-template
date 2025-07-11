# go-fiber-template
Go Fiber project template

- [Fiber v2](https://gofiber.io/)
- [Air](https://github.com/cosmtrek/air) (Hot Reload)
- [GoDotEnv](github.com/joho/godotenv)

## Features

- Fiber static page
- Env vars
- Docker Compose with Postgres
- Database creation on startup
- Model example
- Group routes to api
- Handler to CRUD operations
- Fiber Basic Auth Middleware
- Unit Test example
- Default Response Structure

## Requirements

- Download and install Go 1.23.1 or higher

## Run Database

```shell
docker compose up
```

## Run Application

Check if you have .env file. If not, run:
```shell
cp .env.example .env
```

Run with Air Hot Reload
```shell
air
```

Or run directly on the source
```shell
go run main.go
```

Or compile a binary and run
```shell
go build
./go-fiber-template
```

## Access

- Hello World: http://localhost:3000/
- Static: http://localhost:3000/home
- Api: http://localhost:3000/api

## Tests

Run tests
```shell
go test -v ./...
```

Check test coverage
```shell
go test ./... -cover
```

Generate detailed report
```shell
go test ./... -coverprofile=coverage.out
```

Generate detailed HTML report with source
```shell
go tool cover -html=coverage.out
```
