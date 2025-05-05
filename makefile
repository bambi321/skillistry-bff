project_entry = ./src
project_controller = ./controller
project_handlers = ./handlers
project_helpers = ./helpers
project_services = ./services

go_to_project = cd $(project_entry)

default: tidy start

tidy: 
	go mod tidy

start:
	$(go_to_project) && go run $(project_controller)

start-v2:
	echo "Spinning up API"
	docker compose up --build skillistry-bff

test:
	echo "Running unit tests"
	docker compose run -T \
		--rm \
		golang \
		bash scripts/unit-test.sh $(project_handlers) $(project_helpers) $(project_services)