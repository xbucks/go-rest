package handlers

import (
	"testing"

	"github.com/rameshsunkara/go-rest-api-example/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestNewSeedHandler(t *testing.T) {
	ohandler := NewSeedHandler(&MockMongoDataBase{})

	assert.IsType(t, &SeedHandler{}, ohandler)
	assert.IsType(t, &db.OrdersDataService{}, ohandler.dataSvc)
}
