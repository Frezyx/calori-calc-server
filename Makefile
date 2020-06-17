.PHONY: build
build:
		go build -v ./cmd/apiserver

.PHONY: test
test:
		go test -v -race -timeout 30s ./internal/app/apiserver
		go test -v -race -timeout 30s ./internal/app/store
		go test -v -race -timeout 30s ./internal/app/model

.PHONY: run
run:
		go build -v ./cmd/apiserver
		apiserver

.DEFAULT_GOAL := build 