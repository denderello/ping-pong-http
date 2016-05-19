PROJECT=ping-pong-http
ORGANIZATION=denderello

PROJECT_PATH := "github.com/$(ORGANIZATION)/$(PROJECT)"
BIN := $(PROJECT)

.PHONY: install 

$(BIN): 
	go build .

install:
	go install .
