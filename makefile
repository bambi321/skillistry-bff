project_entry = controller.go
project_handlers = ./handlers
project_helpers = ./helpers
project_services = ./services

go_to_project = cd src

default: tidy start

tidy: 
	$(go_to_project) && go mod tidy

start:
	$(go_to_project) && go run $(project_entry)

test:
	$(go_to_project) && go test -timeout 30s $(project_handlers)
	$(go_to_project) && go test -timeout 30s $(project_helpers)
	$(go_to_project) && go test -timeout 30s $(project_services)