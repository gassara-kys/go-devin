include .env
GOCACHE ?= $(CURDIR)/.gocache
GOLANGCI_LINT_CACHE ?= $(CURDIR)/.golangci-lint-cache

.PHONY: lint
lint:
	GOCACHE=$(GOCACHE) GOLANGCI_LINT_CACHE=$(GOLANGCI_LINT_CACHE) golangci-lint run ./...

.PHONY: test
test:
	GOCACHE=$(GOCACHE) go test ./...

.PHONY: build
build:
	go build ./...

########################################################
## API Test
########################################################
.PHONY: api-session-list
api-session-list:
	@DEVIN_API_KEY=$(DEVIN_API_KEY) go run ./examples/sessions/list
