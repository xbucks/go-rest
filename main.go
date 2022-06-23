package main

import (
	"flag"
	"os"
	"time"

	"github.com/rameshsunkara/go-rest-api-example/internal/server"
	"github.com/rameshsunkara/go-rest-api-example/pkg/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"

	_ "github.com/rameshsunkara/go-rest-api-example/docs"
	"github.com/rameshsunkara/go-rest-api-example/internal/config"
	"github.com/rameshsunkara/go-rest-api-example/internal/db"
	"github.com/rameshsunkara/go-rest-api-example/internal/models"
	"github.com/rs/zerolog/log"
)

const (
	ServiceName = "ecommerce-orders"
	DBName      = "ecommerce"
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
	upTime := time.Now()

	// Parse command line flags
	v := flag.String("version", "0.0", "current version of this service")
	flag.Parse()

	env := os.Getenv("environment")
	if env == "" {
		env = "dev"
	}

	// Metadata of the service
	serviceInfo := &models.ServiceInfo{
		Name:        ServiceName,
		UpTime:      upTime,
		Environment: env,
		Version:     *v,
	}

	// Setup : Log
	setupLog(env)

	log.Log().
		Object("Service", serviceInfo).
		Msg("starting")

	// Load Configuration
	config.LoadConfig(env)

	// Setup : DB
	dbClient, database := db.Init(DBName)

	// Setup : Server
	server.Init(serviceInfo, dbClient, database)

	log.Fatal().Str("ServiceName", ServiceName).Msg("Server Exited")
}

func setupLog(env string) {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	lvl := zerolog.InfoLevel
	logDest := os.Stdout
	logger := zerolog.New(logDest).With().Caller().Timestamp().Logger()
	if util.IsDevMode(env) {
		lvl = zerolog.TraceLevel
		logger = zerolog.New(zerolog.ConsoleWriter{Out: logDest}).With().Caller().Timestamp().Logger()
	}
	zerolog.SetGlobalLevel(lvl)
	log.Logger = logger
}
