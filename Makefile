.DEFAULT_GOAL := help
.PHONY: test test/coverage build docker/build docker/run docker/stop help

test:	## Run test suite
	# go fmt $(go list ./... | grep -v /vendor/)
	# go vet $(go list ./... | grep -v /vendor/)
	go test -v ./...

test/coverage:	## Run test suite with coverage report
	go test -v ./... -cover

build:	## Build project
	go build ./cmd/playground

docker/build: ## Rebuild any images required in docker compose
	docker-compose build -f deployments/docker-compose.yml

docker/run: ## Start running docker compose environment with most recent images
	docker-compose -f deployments/docker-compose.yml up -d

docker/stop: ## Stop running docker compose
	docker-compose -f deployments/docker-compose.yml stop

help:	## Prints this help message
	@grep -E '^[\/a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
