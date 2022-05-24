package main

import (
	"flag"
	"go-rest-api-example/cmd/server"
	_ "go-rest-api-example/docs"
	"go-rest-api-example/internal/config"
	"go-rest-api-example/internal/db"
	customLog "go-rest-api-example/pkg/log"
	"log"
)

// @title           GO Rest Example API Service (Purchase Order Tracker)
// @version         1.0
// @description     A sample service to demonstrate how to develop REST API in golang

// @contact.name    Ramesh Sunkara
// @contact.url
// @contact.email

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	environment := flag.String("environment", "dev", "Specify environment")
	log.Println("Environment:", *environment)
	flag.Parse()
	customLog.SetupLogger(*environment, "go-rest-api-example")
	defer customLog.Logger.Sync()
	config.LoadConfig(*environment)
	db.Init()
	server.Init()
}
