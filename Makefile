build:
	@go build -o bin/main cmd/main.go

run: build
	@./bin/main

run-dev:
	@air -c .air.toml

