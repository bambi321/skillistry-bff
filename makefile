project_entry = controller.go
project_handlers = ./handlers
project_helpers = ./helpers
project_services = ./services

default: tidy start

tidy:
	go mod tidy

start:
	go run $(project_entry)

test:
	go test -timeout 30s $(project_handlers)
	go test -timeout 30s $(project_helpers)
	go test -timeout 30s $(project_services)