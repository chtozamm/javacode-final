.PHONY: run build lint audit test

run:
	./main -c config.env

build:
	GOOS=linux GOARCH=amd64 go build -o main ./cmd

lint:
	go vet ./...
	staticcheck ./...


audit:
	gosec ./...

test:
	go test -vet=off ./...
