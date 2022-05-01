#!make

SHELL=/bin/bash -o pipefail

.PHONY: all test benchmark

all: test benchmark

test:
	go test -coverprofile=coverage.out ./...

benchmark:
	go test -bench=. ./...

coverage:
	go tool cover -html=coverage.out
