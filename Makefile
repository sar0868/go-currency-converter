lint:
	golangci-lint run ./...

lint-fix:
	golangci-lint run ./... --fix

run:
	go run cmd/main.go

test:
	go test ./...
	