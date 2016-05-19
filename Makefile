PROJECT=ping-pong-http
ORGANIZATION=denderello

PROJECT_PATH := "github.com/$(ORGANIZATION)/$(PROJECT)"
BIN := $(PROJECT)
SOURCE=$(shell find . -name '*.go')

.PHONY: install 

$(BIN): $(SOURCE)
	go build .

install:
	go install .
