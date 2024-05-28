# Makefile

.PHONY: server build

# Target to run the Go server
server:
	go run cmd/main.go

# Target to build the binary
build:
	go build -o your_go_binary_name cmd/main.go
