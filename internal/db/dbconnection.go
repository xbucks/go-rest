package db

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/rameshsunkara/go-rest-api-example/internal/config"
	"github.com/rameshsunkara/go-rest-api-example/pkg/log"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectOnce sync.Once
var dbClient *mongo.Client

func Init() {
	c := config.GetConfig()
	connectionUrl := c.GetString("db.dsn")
	log.Logger.Debug("Connection Url:" + connectionUrl)
	connectOnce.Do(func() {
		dbClient = newConnection(connectionUrl)
	})
}

func newConnection(connectionUrl string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(connectionUrl)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Logger.Fatal("Connection Failed to Database", zap.Error(err))
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	connectionError := client.Connect(ctx)
	if connectionError != nil {
		log.Logger.Fatal("Connection Failed to Database", zap.Error(connectionError))
	}

	pingError := client.Ping(context.TODO(), nil)
	if pingError != nil {
		log.Logger.Fatal("Connection Failed to Database", zap.Error(connectionError))
	}
	color.Green("Connection established")
	return client
}

func GetDBClient() *mongo.Client {
	log.Logger.Debug("In GetDB Client")
	return dbClient
}

func Ping() error {
	if dbClient == nil {
		return errors.New("invalid state, you should never be here in an ideal world")
	}
	if err := dbClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		return err
	}
	log.Logger.Debug("Pinged DB successfully")
	return nil
}
