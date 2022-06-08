package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllSuccess(t *testing.T) {
	dSvc := NewOrderDataService(TestDataBase)
	results, _ := dSvc.GetAll()
	assert.EqualValues(t, len(*results), 100)
}

func TestGetByIdSuccess_NoData(t *testing.T) {
	dSvc := NewOrderDataService(TestDataBase)
	const id = "hola-non-id"
	result, _ := dSvc.GetById(id)
	assert.Nil(t, result)
}
