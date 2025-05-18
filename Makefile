VERSION := $(shell git describe --tags --always)
BINARY := smtp2zoho
IMAGE := smtp2zoho

## build          Build the binary with embedded version
build:
	go build -ldflags "-X 'smtp2zoho/config.Version=$(VERSION)'" -o $(BINARY)

## release        Build the Docker image with version tag
release:
	docker build \
		--build-arg VERSION=$(VERSION) \
		-t $(IMAGE):$(VERSION) .

## help           Show available commands
help:
	@grep -E '^##' $(MAKEFILE_LIST) | sed -E 's/^## //;s/:.*//g' | awk '{printf "  \033[36m%-20s\033[0m %s\n", $$1, substr($$0, index($$0,$$2))}'
