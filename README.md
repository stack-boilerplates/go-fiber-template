# go-fiber-template
Go Fiber project template

- [Fiber v2](https://gofiber.io/)
- [Air](https://github.com/cosmtrek/air) (Hot Reload)

## Requirements

- Download and install Go 1.23.1 or higher

## Run

Run directly on the source
```shell
go run main.go
```

Compile a binary and run
```shell
go build
./go-fiber-template
```

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
