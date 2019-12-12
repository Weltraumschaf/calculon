all: build

clean:
	@echo "Deleting files ..."
	@rm -rfv bin

build:
	@echo "Building project ..."
	@mkdir -p bin
	go build -o bin/calculon cmd/calculon/main.go

run:
	@echo "Runing project ..."
	go run cmd/calculon/main.go

test:
	@echo "Testing project ..."
	go test -v ./...

help:
	@echo "Execute one of these targets:"
	@echo " - clean"
	@echo " - build"
	@echo " - run"
	@echo " - test"