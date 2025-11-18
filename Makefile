SHELL := /bin/bash
GOCACHE ?= $(CURDIR)/.gocache

.PHONY: lint test build

lint:
	golangci-lint run ./...

test:
	GOCACHE=$(GOCACHE) go test ./...

build:
	go build ./...
