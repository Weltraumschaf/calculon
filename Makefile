all: build

clean:
	@echo "Deleting files ..."
	@rm -rfv bin

build:
	@echo "Building project ..."
	@mkdir -p bin
	go build -o bin/calculon main.go

test:
	@echo "Testing project ..."
	go test -v ./...

help:
	@echo "Execute one of these targets:"
	@echo " - clean"
	@echo " - build"
	@echo " - run"