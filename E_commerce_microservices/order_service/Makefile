# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run

# Directories
MAIN_DIR=cmd/main.go

# Executable name
BINARY_NAME=myapp

# Targets
all: test build run

build:
	@echo "Building the binary..."
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN_DIR)

clean:
	@echo "Cleaning up..."
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

test:
	@echo "Running tests..."
	$(GOTEST) ./...

run:
	@echo "Running the application..."
	$(GORUN) $(MAIN_DIR)

deps:
	@echo "Getting dependencies..."
	$(GOGET) ./...

.PHONY: all build clean test run deps


