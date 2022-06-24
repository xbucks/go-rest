-include .env

PROJECT_NAME := $(shell basename "$(PWD)")

# GIT commit id will be used as version
VERSION ?= $(shell git rev-parse --short HEAD)

DOCKER_IMAGE_NAME := "$(PROJECT_NAME):$(VERSION)"
DOCKER_CONTAINER_NAME := "$(PROJECT_NAME)-$(VERSION)"

MODULE = $(shell go list -m)

PACKAGES := $(shell go list ./... | grep -v /vendor/)
LDFLAGS := -ldflags "-X main.version=${VERSION}"

CONFIG_FILE ?= ./config/dev.yaml

# PID file will keep the process id of the server in Development Mode.
PID_FILE := $(shell pwd)/tmp/$(PROJECT_NAME).pid

FSWATCH_FILE := './fswatch.cfg'

## start: Starts everything that is required to serve the APIs
start:
	docker-compose up -d
	make run

## develop: Starts API Server in live reload mode and starts the required supplementary services in the background
develop:
	docker-compose up -d
	make run-live

## run: Run the API server alone in normal mode (without supplemantary services such as DB etc.,)
run:
	@export environment=$(profile)
	go run ${LDFLAGS} main.go -version="${VERSION}" & echo $$! >> $(PID_FILE)

## restart: Restarts the API server
restart:
	@pkill -P `cat $(PID_FILE)` || true
	@printf '%*s\n' "80" '' | tr ' ' -
	@echo "Restarting server..."
	make run
	@printf '%*s\n' "80" '' | tr ' ' -

## run-live: Run the API server with live reload support (requires fswatch)
run-live:
	make run
	@fswatch -x -o --event Created --event Updated --event Renamed -r internal pkg config | xargs -n1 -I {} make restart

stop:
	docker-compose down
	@pkill -P `cat $(PID_FILE)` || true

## build: Build the API server binary
build:
	CGO_ENABLED=0 go build ${LDFLAGS} -a -o ${PROJECT_NAME} $(MODULE)

## docker-build: Build the API server as a docker image
docker-build:
	$(info ---> Building Docker Image: ${DOCKER_IMAGE_NAME}, Exposed Port: ${port})
	docker build -t ${DOCKER_IMAGE_NAME} . \
		--build-arg port=${port} \

## docker-run: Run the API server as a docker container
docker-run:
	$(info ---> Running Docker Container: ${DOCKER_CONTAINER_NAME} in Environment: ${profile})
	docker run --name ${DOCKER_CONTAINER_NAME} -it \
				--env environment=${profile} \
				$(DOCKER_IMAGE_NAME)

## docker-start: Builts Docker image and runs it.
docker-start: build-docker run-docker

## docker-stop: Stops the docker container
docker-stop:
	docker stop $(DOCKER_CONTAINER_NAME)

## docker-remove: Removes the docker images and containers	
docker-remove:
	docker rm $(DOCKER_CONTAINER_NAME)
	docker rmi $(DOCKER_IMAGE_NAME)

## docker-clean: Cleans all docker resources
docker-clean: docker-clean-service-images docker-clean-build-images

## docker-clean-service-images: Stops and Removes the service images
docker-clean-service-images: docker-stop docker-remove

## docker-clean-build-images: Removes build images
docker-clean-build-images: 
	docker rmi $(docker images --filter label="builder=true")

## version: Display the current version of the API server
version:
	@echo $(VERSION)

## api-docs: Generate OpenAPI3 Spec
api-docs:
	swag init -g main.go
	curl -X POST "https://converter.swagger.io/api/convert" \
		-H "accept: application/json" \
		-H "Content-Type: application/json" \
		-d @docs/swagger.json > docs/openapi.json

## test: Run tests
test:
	go test -v ./...

## coverage: Measures code coverage
coverage:
	go test -race -coverprofile=coverage.out -covermode=atomic ./...

.PHONY: help
help: Makefile
	@echo
	@echo " Choose a command to run in "$(PROJECT_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |	sed -e 's/^/ /'
	@echo