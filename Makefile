run:
	swag init
	golangci-lint run
	go mod tidy
	go run main.go