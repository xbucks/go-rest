[![Build Status](https://github.com/rameshsunkara/go-rest-api-example/actions/workflows/cibuild.yml/badge.svg)](https://github.com/rameshsunkara/go-rest-api-example/actions/workflows/cibuild.yml?query=+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/rameshsunkara/go-rest-api-example)](https://goreportcard.com/report/github.com/rameshsunkara/go-rest-api-example)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/rameshsunkara/go-rest-api-example)](https://www.tickgit.com/browse?repo=github.com/rameshsunkara/go-rest-api-example)
[![Security Analysis](https://github.com/rameshsunkara/go-rest-api-example/actions/workflows/codeql.yml/badge.svg)](https://github.com/rameshsunkara/go-rest-api-example/actions/?query=workflow%3ACodeQL+branch%3Amain)

# REST API microservice in golang

## Why?

There are many open source boilerplate repos but why I did this ?

1. Coming from years of building Full Stack application in ReactJS and JVM based languages, I did not like any of them
   completely.
   So I created my own while obeying 'GO' principles and guidelines.
   You will find a lot of similarities in this repo when compared to the most popular go boilerplate templates because I
   probably borrowed
   ideas from them. (my apologies if I failed to miss any of them in the references)

2. I want to pick the tools for Routing, Logging, Configuration Management etc., to my liking and preferences.

3. I want a version which I have full control to change/update based on my professional work requirements.

### QuickStart

Pre-requisites:

- Docker Compose

1. Start the service

        make start

2. Visit Swagger API in browser

        http://localhost:8080

   If you are a fan of Postman, import the included [Postman collection](Orders.postman_collection.json) which is actually a better option as DB is not
   seeded yet.

### QuickStart - Develop

Pre-requisites: Docker, Docker Compose, [Swag](https://github.com/swaggo/swag), fswatch

1. Start the service in live reload mode

        make develop

2. Use postman to explore the existing APIs, start with SeedDB request.

Other Options:

Choose a command to run in go-rest-api-example:

      start                                      Starts everything that is required to serve the APIs
      develop                                    Starts API Server in live reload mode and starts the required supplementary services in the background
      run                                        Run the API server alone in normal mode (without supplemantary services such as DB etc.,)
      restart                                    Restarts the API server
      run-live                                   Run the API server with live reload support (requires fswatch)
      build                                      Build the API server binary
      build-docker                               Build the API server as a docker image
      run-docker                                 Run the API server as a docker container
      docker-start                               Builts Docker image and runs it.
      docker-stop                                Stops the docker container
      docker-remove                              Removes the docker images and containers   
      docker-clean                               Cleans all docker resources
      docker-clean-service-images                Stops and Removes the service images
      docker-clean-build-images                  Removes build images
      version                                    Display the current version of the API server
      lint                                       Runs golint on all Go packages (TODO)
      fmt                                        Run format "go fmt" on all Go packages
      api-docs                                   Generate OpenAPI3 Spec

### Tools

1. Routing - [Gin](https://github.com/gin-gonic/gin)
2. Logging - [Zap](https://github.com/uber-go/zap)
3. Configuration - [Viper](https://github.com/spf13/viper)
4. Database - [Mongo](https://www.mongodb.com/)
5. Container - [Docker](https://www.docker.com/)
6. API Spec Generation - [Swag](https://github.com/swaggo/swag)

### Features

- Live reload
- OpenApi3 Spec generation
- Log file rotation
- Easy to use 'make' tasks
- Multi-Stage container build (cache enabled)
- Tag docker images with latest git commit

### TODO

- [ ] Add more and clear documentation about the features this offers and how to replace tools
- [ ] Automate Open API3 Spec Generation completely
- [ ] Seed local DB through docker and add DB Migration Support
- [ ] Add more unit tests
- [ ] Add more profiles and obey all [12-Factor App rules](https://12factor.net/ru/)
- [ ] Add CI/CD tooling & Necessary Github hooks
- [ ] Add missing references/inspirations
- [ ] Improvements to the api in terms of error handling, proper messaging etc., ( that wasn't focus)
- [ ] API Documentation - Lot of potential to improve

### References

- [gin-boilerplate](https://github.com/Massad/gin-boilerplate)
- [go-rest-api](https://github.com/qiangxue/go-rest-api)
- [go-base](https://github.com/dhax/go-base)

### Contribute

- Please feel free to Open PRs
- Please create issues with any problem you noticed

## Known Issues

- Default Swagger Docs doesn't use generated OpenAPI3 Spec
- Docker Run make targets are for build server for now
- LintCI should be fixed
