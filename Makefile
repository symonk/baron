## Variables
GOCMD=go
BINARY_NAME=baron


all: help

## Build:
build:  ## Build the project and put the output binary in dist/binary*
	GO111MODULE=on $(GOCMD) build -o dist/binary/$(BINARY_NAME) ./cmd/baron

clean:  ## Remove build related files
	rm -fr ./dist/binary/*

help:  ## Show this help.
	@echo work in progress
