build:
	go build ./cmd/gendiff

test:
	go test ./...

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8 run

lint-fix:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8 run --fix
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8 fmt
