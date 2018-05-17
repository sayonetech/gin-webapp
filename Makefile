# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=go-webapp
BINARY_NAME_MIGRATE=go-webapp-migrate
BINARY_UNIX=$(BINARY_NAME)_unix

all: run

build:
	$(GOBUILD) -o ./build/$(BINARY_NAME) -v ./

test:
	govendor test +local

clean:
	$(GOCLEAN)
	rm -f ./build/$(BINARY_NAME)
	rm -f ./build/$(BINARY_UNIX)

run:
	$(GOBUILD) -o ./build/$(BINARY_NAME) -v ./
	./build/$(BINARY_NAME)

runlocal:
	echo $(SHELL)
	. ./env/local.env
	$(MAKE) run

restart:
	kill -INT $$(cat pid)
	$(GOBUILD) -o ./build/$(BINARY_NAME) -v ./
	./build/$(BINARY_NAME)

deps:
	$(GOGET) github.com/kardianos/govendor
	govendor sync

cross:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./build/$(BINARY_NAME) -v ./


migrate:
	$(GOBUILD) -o ./build/$(BINARY_NAME_MIGRATE) -v ./migration/migration.go
	./build/$(BINARY_NAME_MIGRATE)


loadfixtures:
	#TODO
