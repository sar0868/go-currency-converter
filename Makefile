lint:
	golangci-lint run ./...

lint-fix:
	golangci-lint run ./... --fix

run:
	go run ./...

test:
	go test ./...
	