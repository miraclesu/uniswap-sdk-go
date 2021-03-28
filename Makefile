GO ?= go
GOFMT ?= gofmt "-s"
VERSION ?= $(shell git describe --tags --always || git rev-parse --short HEAD)
GOFILES := $(shell find . -name "*.go" -type f -not -path "./vendor/*")

.PHONY: all
all: test

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	$(GO) test -race -coverprofile=coverage.txt -covermode=atomic ./...

version:
	@echo $(VERSION)
