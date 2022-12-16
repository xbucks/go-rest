package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectOnce sync.Once

// ConnectionTimeOut - Max time to establish DB connection // TODO: Move to config
const ConnectionTimeOut = 10 * time.Second

type MongoDatabase interface {
	Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection
}

type MongoManager interface {
	Database() MongoDatabase
	Ping() error
	Disconnect() error
}

// connectionManager - Implements MongoManager
type connectionManager struct {
	client   *mongo.Client
	database *mongo.Database
}

// NewMongoManager - Initializes DB connection and returns a Manager object which can be used to perform DB operations
func NewMongoManager(dbName string, connUrl string) (MongoManager, error) {
	log.Debug().Str("DB Connection Url", connUrl)

	dbMgr := &connectionManager{}
	var connErr error
	connectOnce.Do(func() {
		if c, err := newConnection(connUrl); err != nil {
			connErr = err
		} else {
			db := c.Database(dbName)
			dbMgr.database = db
			dbMgr.client = c

			// Verify connection
			if err := dbMgr.Ping(); err != nil {
				connErr = err
			}
		}
	})

	return dbMgr, connErr
}

// newConnection - Establishes connection using given connection url and returns mongo client
func newConnection(connectionUrl string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(connectionUrl)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Error().Err(err).Msg("Connection Failed to Database")
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), ConnectionTimeOut)
	defer cancel()
	connErr := client.Connect(ctx)
	if connErr != nil {
		log.Error().Err(connErr).Msg("Connection Failed to Database")
		return nil, err
	}

	return client, nil
}

// Database - Returns configured database instance
func (c *connectionManager) Database() MongoDatabase {
	return c.database
}

// Ping - Validates application's connectivity to the underlying database by pinging
func (c *connectionManager) Ping() error {
	if err := c.client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Error().Err(err).Msg("unable to connect to DB")
		return err
	}
	return nil
}

// Disconnect - Close connection to Database
func (c *connectionManager) Disconnect() error {
	log.Info().Msg("Disconnecting from Database")
	if err := c.client.Disconnect(context.Background()); err != nil {
		log.Error().Err(err).Msg("unable to disconnect from DB")
		return err
	}
	log.Info().Msg("Successfully disconnected from DB")
	return nil
}
