all: build

.PHONY: clean
clean:
	@echo "Deleting files ..."
	@rm -rfv bin

.PHONY: build
build:
	@echo "Building project ..."
	@mkdir -p bin
	go build -o bin/calculon cmd/calculon/main.go

.PHONY: run
run:
	@echo "Runing project ..."
	go run cmd/calculon/main.go

.PHONY: test
test:
	@echo "Testing project ..."
	go test -v ./...

.PHONY: help
help:
	@echo "Execute one of these targets:"
	@echo " - clean"
	@echo " - build"
	@echo " - run"
	@echo " - test"