all: build

clean:
	@echo "Deleting files ..."
	@rm -rfv bin

build:
	@echo "Building project ..."
	@mkdir -p bin
	go build -o bin/ipcalc main.go

run:
	@echo "Runing project ..."
	go run main.go

help:
	@echo "Execute one of these targets:"
	@echo " - clean"
	@echo " - build"
	@echo " - run"