BASE_APP_NAME := microwms

ifeq ($(OS), Windows_NT)
	PROJECT_NAME=$(BASE_APP_NAME).exe
else
	PROJECT_NAME=$(BASE_APP_NAME)
endif

VERSION?=$(shell git describe --tags --always --dirty)
COVER_FILE ?= $(PROJECT_NAME)-coverage.out

define setup_env
	$(eval ENV_FILE :=$(1).env)
	@echo "Setting up environment file: $(ENV_FILE)"
	$(eval include $(ENV_FILE))
	$(eval export)
endef

.PHONY: devEnv
devEnv:
	$(call setup_env, dev)

.PHONY: tools
tools: ## Install all needed tools, e.g. for static checks
	@echo Installing tools from tools-versions.txt
	@xargs -tI % go install % < tools-versions.txt

.PHONY: lint
lint: tools ## Check the project with lint
	golangci-lint run -c .golangci.yml --fix ./...

.PHONY: test
test: ## Run unit (short) tests
	go test -short ./... -coverprofile=$(COVER_FILE)
	go tool cover -func=$(COVER_FILE) | grep ^total

.PHONY: check
check: lint test ## Check project with static checks and tests

.PHONY: dep
dep: ## Manage go mod dependencies, beautify go.mod and go.sum files, vendor them for safety
	go mod tidy
	go mod vendor

.PHONY: build
build: ## Build the project binary
	go build -o $(PROJECT_NAME) -ldflags "-X main.version==$(VERSION)" ./cmd/$(BASE_APP_NAME)/

.PHONY: run
run: devEnv ## Start the project
	go run ./cmd/$(BASE_APP_NAME)/main.go


.PHONY: swagger
swagger: ## Build swagger doc
	swag init -g cmd/$(BASE_APP_NAME)/main.go -o docs/swagger --parseVendor

