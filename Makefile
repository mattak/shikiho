GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
BINARY_DIR=bin
BINARY_NAME1=shikiho2json
BINARY_NAME2=shikiho2growthtsv
TARGET_FILE1=./cmd/shikiho2json/main.go
TARGET_FILE2=./cmd/shikiho2growthtsv/main.go

all: test build
install: build system_install

.PHONY: test
test:
	$(GOTEST) -v ./test/shikiho/

.PHONY: build
build:
	$(GOBUILD) -o $(BINARY_DIR)/$(BINARY_NAME1) $(TARGET_FILE1)
	$(GOBUILD) -o $(BINARY_DIR)/$(BINARY_NAME2) $(TARGET_FILE2)

.PHONY: run1
run1:
	$(GORUN) $(TARGET_FILE1)

.PHONY: run2
run2:
	$(GORUN) $(TARGET_FILE2)

.PHONY: clean
clean:
	$(GOCLEAN)
	rm -r $(BINARY_DIR)

.PHONY: system_install
system_install:
	cd cmd/shikiho2json && go install
	cd cmd/shikiho2growthtsv && go install
