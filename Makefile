GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
BINARY_NAME=shikiho2json
BINARY_DIR=bin
TARGET_FILE=./cmd/shikiho2json/main.go

all: test build
install: build system_install

.PHONY: test
test:
	$(GOTEST) -v ./test/shikiho/

.PHONY: build
build:
	$(GOBUILD) -o $(BINARY_DIR)/$(BINARY_NAME) $(TARGET_FILE)

.PHONY: run
run:
	$(GORUN) $(TARGET_FILE)

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -r $(BINARY_DIR)

.PHONY: system_install
system_install:
	cd cmd/shikiho2json && go install
