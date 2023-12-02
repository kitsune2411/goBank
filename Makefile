build:
	@go build -o bin/goBank.exe

run: build
	@./bin/goBank.exe

test: 
	@go test -v ./...