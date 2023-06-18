.PHONY: generate build install

generate:
	go generate ./cli

build:
	go build ./cmd/...

install:
	go install ./cmd/...

