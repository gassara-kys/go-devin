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

.PHONY: api-session-get
api-session-get:
	@DEVIN_API_KEY=$(DEVIN_API_KEY) \
	DEVIN_SESSION_ID=$(DEVIN_SESSION_ID) \
	go run ./examples/sessions/get

.PHONY: api-session-create
api-session-create:
	@DEVIN_API_KEY=$(DEVIN_API_KEY) go run ./examples/sessions/create

.PHONY: api-session-send-message
api-session-send-message:
	@DEVIN_API_KEY=$(DEVIN_API_KEY) \
	DEVIN_SESSION_ID=$(DEVIN_SESSION_ID) \
	go run ./examples/sessions/send_message

.PHONY: api-secret
api-secret:
	@DEVIN_API_KEY=$(DEVIN_API_KEY) go run ./examples/secrets/list

.PHONY: api-playbook
api-playbook:
	@DEVIN_API_KEY=$(DEVIN_API_KEY) go run ./examples/playbooks/list

.PHONY: api-knowledge
api-knowledge:
	@DEVIN_API_KEY=$(DEVIN_API_KEY) go run ./examples/knowledge/list
