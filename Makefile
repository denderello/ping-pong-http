.PHONY: build install 

build:
	go build .

install:
	go install .

test:
	./test.sh
