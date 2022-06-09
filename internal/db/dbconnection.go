package db

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/rameshsunkara/go-rest-api-example/internal/config"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	dbClient    *mongo.Client
	database    *mongo.Database
	connectOnce sync.Once
)

func Init(dbName string) (*mongo.Client, *mongo.Database) {
	c := config.GetConfig()
	connectionUrl := c.GetString("db.dsn")
	log.Debug().Str("Connection Url", connectionUrl)
	connectOnce.Do(func() {
		newConnection(connectionUrl, dbName)
	})

	return dbClient, database
}

func newConnection(connectionUrl string, dbName string) {
	clientOptions := options.Client().ApplyURI(connectionUrl)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal().Err(err).Msg("Connection Failed to Database")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	connErr := client.Connect(ctx)
	if connErr != nil {
		log.Fatal().Err(connErr).Msg("Connection Failed to Database")
	}

	pingError := client.Ping(context.TODO(), nil)
	if pingError != nil {
		log.Fatal().Err(pingError).Msg("Connection Failed to Database")
	}

	// Set the globals
	dbClient = client
	database = client.Database(dbName)
}

func GetDBClient() (*mongo.Client, error) {
	if dbClient == nil {
		return nil, errors.New("invalid state, database.Init is not called")
	}
	return dbClient, nil
}

func GetDB() (*mongo.Database, error) {
	if database == nil {
		return nil, errors.New("invalid state, database.Init is not called")
	}
	return database, nil
}

func OverrideDBSetup(c *mongo.Client, db *mongo.Database) {
	dbClient = c
	database = db
}
