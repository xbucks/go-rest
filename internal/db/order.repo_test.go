package db

import (
	"github.com/bxcodec/faker/v3"
	"github.com/rameshsunkara/go-rest-api-example/internal/models"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"math/rand"
	"testing"
)

var orderId primitive.ObjectID

func TestNewOrderDataService(t *testing.T) {
	d, _ := dbMgr.Database()
	ds := NewOrderDataService(d)
	assert.IsType(t, &OrdersDataService{}, ds)
	assert.IsType(t, &mongo.Collection{}, ds.collection)
	assert.EqualValues(t, ds.collection.Name(), OrdersCollection)
}

func TestCreateSuccess(t *testing.T) {
	d, _ := dbMgr.Database()
	dSvc := NewOrderDataService(d)
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
	result, err := dSvc.Create(po)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to insert data")
	}
	orderId = result.InsertedID.(primitive.ObjectID)
	assert.NotNil(t, result.InsertedID)
}

func TestCreate_InvalidReq(t *testing.T) {
	d, _ := dbMgr.Database()
	dSvc := NewOrderDataService(d)
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
		ID:       primitive.NewObjectID(),
		Products: product,
	}
	_, err := dSvc.Create(po)
	assert.Error(t, err)
}

func TestGetAllSuccess(t *testing.T) {
	d, _ := dbMgr.Database()
	dSvc := NewOrderDataService(d)
	results, _ := dSvc.GetAll()
	orders := results.(*[]models.Order)
	assert.EqualValues(t, 100, len(*orders))
}

/*
func TestGetByIdSuccess(t *testing.T) {
	d, _ := dbMgr.Database()
	dSvc := NewOrderDataService(d)
	result, _ := dSvc.GetById(orderId.String())
	order := result.(models.Order)
	assert.NotNil(t, result)
	assert.EqualValues(t, orderId, order.ID.String())
}
*/

func TestGetByIdSuccess_NoData(t *testing.T) {
	d, _ := dbMgr.Database()
	dSvc := NewOrderDataService(d)
	const id = "hola-non-id"
	result, _ := dSvc.GetById(id)
	assert.Nil(t, result)
}
