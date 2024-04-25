run: build
	@./bin/api

build:
	@go build -o ./bin/api ./cmd/api

test:
	@go test -v ./...
