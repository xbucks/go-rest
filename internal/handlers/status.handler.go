package handlers

import (
	"context"
	"github.com/rameshsunkara/go-rest-api-example/internal/models"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/gin-gonic/gin"
)

type ServiceStatus string

const (
	UP   ServiceStatus = "ok"
	DOWN ServiceStatus = "down"
)

type StatusResponse struct {
	Status      ServiceStatus
	ServiceName string
	UpTime      time.Time
	Environment string
	Version     string
}

// MongoDBClient - Enables mocking
type MongoDBClient interface {
	Ping(ctx context.Context, rp *readpref.ReadPref) error
}

type StatusHandler struct {
	svcInfo  *models.ServiceInfo
	dbClient MongoDBClient
}

func NewStatusHandler(s *models.ServiceInfo, client MongoDBClient) *StatusHandler {
	return &StatusHandler{
		svcInfo:  s,
		dbClient: client,
	}
}

func (sc *StatusHandler) CheckStatus(c *gin.Context) {
	log.Debug().Msg("in CheckStatus")
	var stat ServiceStatus
	var code int
	err := sc.dbClient.Ping(context.TODO(), readpref.Primary())
	if err == nil {
		stat = UP
		code = http.StatusOK
	} else {
		log.Error().Msg("unable to connect to DB")
		stat = DOWN
		code = http.StatusFailedDependency
	}
	status := &StatusResponse{
		Status:      stat,
		ServiceName: sc.svcInfo.Name,
		UpTime:      sc.svcInfo.UpTime,
		Environment: sc.svcInfo.Environment,
		Version:     sc.svcInfo.Version,
	}
	c.JSON(code, status)
}
