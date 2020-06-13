.PHONY: build
build:
		go build -v ./cmd/apiserver

.PHONY: test
test:
		go test -v -race -timeout 30s ./internal/app/apiserver
		go test -v -race -timeout 30s ./internal/app/store
		go test -v -race -timeout 30s ./internal/app/model

.DEFAULT_GOAL := build 