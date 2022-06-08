package db

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/rameshsunkara/go-rest-api-example/internal/models"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/strikesecurity/strikememongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	TestDataBase *mongo.Database
)

func TestMain(m *testing.M) {
	mongoServer, err := strikememongo.Start("4.0.5") // TODO: Only this version works, figure out why ?
	if err != nil {
		log.Fatal().Err(err)
	}
	defer mongoServer.Stop()

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoServer.URI()))
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal().Err(err)
	}
	TestDataBase = client.Database(strikememongo.RandomDatabase())
	OverrideDBSetup(client, TestDataBase)
	insertTestData()

	os.Exit(m.Run())
}

func insertTestData() {
	db, err := GetDB()
	if err != nil {
		log.Panic().Err(err).Msg("database is not initialized")
	}
	dSvc := NewOrderDataService(db)

	for i := 0; i < 500; i++ {
		product := []models.Product{
			{
				Name:      faker.Name(),
				Price:     (uint)(rand.Intn(90) + 10),
				Remarks:   faker.Sentence(),
				UpdatedAt: faker.TimeString(),
			},
			{
				Name:      faker.Name(),
				Price:     (uint)(rand.Intn(1000) + 10),
				Remarks:   faker.Sentence(),
				UpdatedAt: faker.TimeString(),
			},
		}

		po := &models.Order{
			Products: product,
		}
		_, err := dSvc.Create(po)
		if err != nil {
			log.Fatal().Err(err).Msg("unable to insert data")
		}
	}
}

func TestDBSuccess(t *testing.T) {
	db, err := GetDB()
	assert.Nil(t, err)
	assert.NotNil(t, db)
	assert.IsType(t, *db, mongo.Database{})
}

func TestDBClient(t *testing.T) {
	client, err := GetDBClient()
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.IsType(t, *client, mongo.Client{})
}
