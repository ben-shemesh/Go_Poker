build:
	@go build -o bin/GoPoker
run: build
 @./bin/GoPoker
 test:
	@go test -v ./...