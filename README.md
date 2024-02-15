# REST API microservice in golang

## Why?

There are many open source boilerplate repos but why I did this ?

1. Coming from years of building Full Stack application in ReactJS and JVM based languages, I did not like any of them
   completely.
   So I created my own while obeying 'GO' principles and guidelines.
   You will find a lot of similarities in this repo when compared to the most popular go boilerplate templates because I
   probably borrowed ideas from them. (my apologies if I failed to miss any of them in the references)

2. I wanted to pick the tools for Routing, Logging, Configuration Management etc., to my liking and preferences.

3. I wanted a version where I have full control to change/update based on my professional work requirements.

### QuickStart

Pre-requisites: Docker, Docker Compose, [Swag](https://github.com/swaggo/swag)

1. Start the service

        make start

2. Visit Swagger API in browser

        open http://localhost:8080/swagger/index.html

   If you are a fan of Postman, import the included [Postman collection](orders.postman_collection.json)

Other Options:

Choose a command to run in go-rest-api-example:

        start                         Starts everything that is required to serve the APIs
        run                           Run the API server alone in normal mode (without supplemantary services such as DB etc.,)
        build                         Build the API server binary
        docker-build                  Build the API server as a docker image
        docker-run                    (WIP) Run the API server as a docker container
        docker-start                  Builts Docker image and runs it.
        docker-stop                   Stops the docker container
        docker-remove                 Removes the docker images and containers
        docker-clean                  Cleans all docker resources
        docker-clean-service-images   Stops and Removes the service images
        docker-clean-build-images     Removes build images
        version                       Display the current version of the API server
        api-docs                      Generate OpenAPI3 Spec
        test                          Run tests
        coverage                      Measures code coverage

### Tools

1. Routing - [Gin](https://github.com/gin-gonic/gin)
2. Logging - [zerolog](https://github.com/rs/zerolog)
3. Configuration - [Viper](https://github.com/spf13/viper)
4. Database - [Mongo](https://www.mongodb.com/)
5. Container - [Docker](https://www.docker.com/)
6. API Spec Generation - [Swag](https://github.com/swaggo/swag)

### Features

- OpenApi3 Spec generation
- Easy to use 'make' tasks to do everything
- Multi-Stage container build (cache enabled)
- Versioning using git commit (both Application and Docker objects)
- Git Actions to build, security analysis and to run code coverage
- Templated Docker and Make files

### TODO

- [ ] Add more and clear documentation about the features this offers and how to replace tools
- [ ] Automate Open API3 Spec Generation completely
- [ ] Add DB Migration Support
- [ ] Add more profiles and obey all [12-Factor App rules](https://12factor.net/ru/)
- [ ] Deploy to cloud
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
