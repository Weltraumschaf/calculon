# Disable built-in rules and variables because we do not need them.
# - https://www.gnu.org/software/make/manual/html_node/Catalogue-of-Rules.html#Catalogue-of-Rules
# - https://www.gnu.org/software/make/manual/html_node/Implicit-Variables.html#Implicit-Variables
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-builtin-variables

PROJECT_DIR 	= $(shell pwd)
BIN_DIR			= $(PROJECT_DIR)/bin
TOOLS_DIR		= $(PROJECT_DIR)/tools
COVERAGE_DIR 	= $(PROJECT_DIR)/coverage

all: build

.PHONY: clean ## Clean the project.
clean:
	rm -rfv $(BIN_DIR) $(COVERAGE_DIR)

.PHONY: build
build: ## Build the project
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/calculon $(PROJECT_DIR)/cmd/calculon/main.go
	@echo "Binary is located in $(BIN_DIR)"

.PHONY: run
run: ## Run the project's binary.
	go run $(PROJECT_DIR)/cmd/calculon/main.go 192.168.123.5/24

.PHONY: test
test: ## Run the unit tests.
	go test -v ./...

cover: ## Generate code coverage report
	$(TOOLS_DIR)/coverage.sh $(COVERAGE_DIR)

cover-html: ## Generate HTML code coverage report
	$(TOOLS_DIR)/coverage.sh $(COVERAGE_DIR) html;

.PHONY: help
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
