.PHONY: all
all: test

.PHONY: format
format:
	go fmt ./...

.PHONY: vet
vet: format
	go vet ./...

.PHONY: test
test: vet
	go test ./...

GOLANGCI_LINT ?= ./bin/golangci-lint
golangci-lint:
ifeq (, $(shell which ./bin/golangci-lint 2> /dev/null))
	@{ \
	set -e ;\
	VERSION="v1.42.1" ;\
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/$${VERSION}/install.sh | sh -s -- -b ./bin $${VERSION} ;\
	}
endif

.PHONY: lint
lint: golangci-lint
	$(GOLANGCI_LINT) run $(LINT_OPTIONS) --verbose --deadline 10m ./...
