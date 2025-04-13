project_entry = controller.go

default: tidy start

tidy:
	go mod tidy

start:
	go run $(project_entry)