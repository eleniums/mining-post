EXECUTABLE=miningpost
WINDOWS=$(EXECUTABLE)_windows_amd64.exe
LINUX=$(EXECUTABLE)_linux_amd64
DARWIN=$(EXECUTABLE)_darwin_amd64
VERSION=$(shell git describe --tags --always --long --dirty)

.PHONY: all test release clean

all: test build ## Build and run tests

test: ## Run unit tests
	./scripts/test_unit.sh

docker: linux ## Build Docker container
	docker build -t $(EXECUTABLE) .

build: windows linux darwin ## Build binaries
	@echo version: $(VERSION)

windows: $(WINDOWS) ## Build for Windows

linux: $(LINUX) ## Build for Linux

darwin: $(DARWIN) ## Build for Darwin (macOS)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 go build -o $(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)" ./cmd/main.go

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -o $(LINUX) -ldflags="-s -w -X main.version=$(VERSION)" ./cmd/main.go

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -o $(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)" ./cmd/main.go

release: ## Create release
	./scripts/release.sh

clean: ## Remove previous build
	rm -f $(WINDOWS) $(LINUX) $(DARWIN)

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'