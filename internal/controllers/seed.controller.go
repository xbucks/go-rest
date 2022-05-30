package controllers

import (
	"go-rest-api-example/internal/db"
	"go-rest-api-example/internal/models"
	"math/rand"
	"net/http"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
)

const (
	SEED_RECORD_COUNT = 50
)

type SeedDBController struct {
	DBService db.OrdersCrudService
}

func (seedController *SeedDBController) SeedDB(c *gin.Context) {
	for i := 0; i < SEED_RECORD_COUNT; i++ {
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
		_, err := seedController.DBService.CreateOrder(po)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Unable inserted data",
			})
			panic("Unable to insert data")
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully inserted fake data",
		"Count":   SEED_RECORD_COUNT,
	})
}
