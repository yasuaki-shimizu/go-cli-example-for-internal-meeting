.PHONY: build test test-cov

build:
	go build -o yakiniku.exe main.go

test:
	go test -v ./...

test-cov:
	go test -cover -coverprofile=cover.out -v ./... && go tool cover -html=cover.out