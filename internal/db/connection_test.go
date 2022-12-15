package db_test

import (
	"context"
	"math/rand"
	"os"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/rameshsunkara/go-rest-api-example/internal/db"
	"github.com/rameshsunkara/go-rest-api-example/internal/models"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/strikesecurity/strikememongo"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	testDBMgr db.MongoManager
)

func TestMain(m *testing.M) {
	mongoServer, err := strikememongo.Start("4.0.5") // TODO: Only this version works, figure out why ?
	if err != nil {
		log.Fatal().Err(err)
	}
	defer mongoServer.Stop()

	d, dErr := db.Init(strikememongo.RandomDatabase(), mongoServer.URI())
	if dErr != nil {
		log.Fatal().Err(dErr)
	}
	testDBMgr = d
	insertTestData()

	os.Exit(m.Run())
}

func insertTestData() {
	database, err := testDBMgr.Database()
	if err != nil {
		log.Panic().Err(err).Msg("database is not initialized")
	}
	dSvc := db.NewOrderDataService(database)

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
		_, err := dSvc.Create(context.TODO(), po)
		if err != nil {
			log.Fatal().Err(err).Msg("unable to insert data")
		}
	}
}

func TestDatabase(t *testing.T) {
	d, err := testDBMgr.Database()
	assert.Nil(t, err)
	assert.NotNil(t, d)
	assert.IsType(t, &mongo.Database{}, d)
}

func TestPing(t *testing.T) {
	err := testDBMgr.Ping()
	assert.Nil(t, err)
}
