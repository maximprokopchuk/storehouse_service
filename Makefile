.PHONY: build

build: 
		go build -v ./cmd/storehouse_service

.PHONY: run

run: 
		go run -v ./cmd/storehouse_service

.PHONY: test
test:
		go test -v -race -timeout 30s ./...

.PHONY: proto_generate
proto_generate:
		protoc -I proto --go_out=plugins=grpc:pkg/api proto/storehouse.proto

.PHONY: sqlc_generate
sqlc_generate:
		sqlc generate

.PHONY: deps
deps:
		go mod tidy

.PHONY: lint
lint:
		golangci-lint run


DEFAULT_GOAL := build