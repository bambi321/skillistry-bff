default: tidy start

tidy:
	go mod tidy

start:
	go run src/controller.go