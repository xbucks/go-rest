package db

import (
	"math/rand"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/rameshsunkara/go-rest-api-example/internal/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
		t.Fail()
	}
	orderId = result.InsertedID.(primitive.ObjectID)
	assert.True(t, !orderId.IsZero())
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

func TestUpdateSuccess(t *testing.T) {
	d, _ := dbMgr.Database()
	dSvc := NewOrderDataService(d)
	product := []models.Product{
		{
			Name:      faker.Name(),
			Price:     (uint)(rand.Intn(90) + 10),
			Remarks:   faker.Sentence(),
			UpdatedAt: faker.TimeString(),
		},
	}

	po := &models.Order{
		ID:       orderId,
		Products: product,
	}
	result, err := dSvc.Update(po)
	assert.EqualValues(t, 1, result)
	assert.Nil(t, err)
}

func TestUpdate_InvalidId(t *testing.T) {
	d, _ := dbMgr.Database()
	dSvc := NewOrderDataService(d)
	product := []models.Product{
		{
			Name:      faker.Name(),
			Price:     (uint)(rand.Intn(90) + 10),
			Remarks:   faker.Sentence(),
			UpdatedAt: faker.TimeString(),
		},
	}

	po := &models.Order{
		ID:       primitive.NilObjectID,
		Products: product,
	}
	result, err := dSvc.Update(po)
	assert.EqualValues(t, 0, result)
	assert.Error(t, err)
}

func TestGetAllSuccess(t *testing.T) {
	d, _ := dbMgr.Database()
	dSvc := NewOrderDataService(d)
	results, _ := dSvc.GetAll()
	orders := results.(*[]models.Order)
	assert.EqualValues(t, 100, len(*orders))
}

func TestGetByIdSuccess(t *testing.T) {
	d, _ := dbMgr.Database()
	dSvc := NewOrderDataService(d)
	result, _ := dSvc.GetById(orderId.Hex())
	order := result.(*models.Order)
	assert.NotNil(t, result)
	assert.EqualValues(t, orderId, order.ID)
}

func TestGetByIdSuccess_NoData(t *testing.T) {
	d, _ := dbMgr.Database()
	dSvc := NewOrderDataService(d)
	const id = "000000000000000000000000"
	result, err := dSvc.GetById(id)
	assert.Nil(t, result)
	assert.Nil(t, err)
}

func TestGetById_InvalidId(t *testing.T) {
	d, _ := dbMgr.Database()
	dSvc := NewOrderDataService(d)
	result, err := dSvc.GetById("i-am-an-invalid-id")
	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestDeleteByIdSuccess(t *testing.T) {
	d, _ := dbMgr.Database()
	dSvc := NewOrderDataService(d)
	result, err := dSvc.DeleteById(orderId.Hex())
	assert.Nil(t, err)
	assert.EqualValues(t, 1, result)
}

func TestDeleteByIdSuccess_NoData(t *testing.T) {
	d, _ := dbMgr.Database()
	dSvc := NewOrderDataService(d)
	const id = "000000000000000000000000"
	result, err := dSvc.DeleteById(id)
	assert.Nil(t, err)
	assert.EqualValues(t, 0, result)
}

func TestDeleteById_InvalidId(t *testing.T) {
	d, _ := dbMgr.Database()
	dSvc := NewOrderDataService(d)
	result, err := dSvc.DeleteById("i-am-an-invalid-id")
	assert.EqualValues(t, 0, result)
	assert.Error(t, err)
}
