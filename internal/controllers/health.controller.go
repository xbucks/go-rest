package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rameshsunkara/go-rest-api-example/internal/db"
	"github.com/rameshsunkara/go-rest-api-example/internal/models"
	"github.com/rameshsunkara/go-rest-api-example/pkg/log"
)

type HealthController struct {
	ServiceName string
	UpTime      time.Time
	Environment string
	Version     string
}

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

func NewHealthController(sInfo *models.ServiceMeta) *HealthController {
	s := &HealthController{
		ServiceName: sInfo.Name,
		UpTime:      sInfo.Uptime,
		Environment: sInfo.Environment,
		Version:     sInfo.Version,
	}
	return s
}

// Status godoc
// @Summary      Status/Health of the service (Doesn't work in Swagger UI)
// @Description  Provides the current status/health of the service. It checks the health of all the critical components of the service.
// @Tags         Status
// @Accept       json
// @Produce      plain
// @Success      200  {string}
// @Failure      424  {string}  string  "down"
// @Router       /health [get]
func (h *HealthController) Status(c *gin.Context) {
	log.Logger.Debug("In Status Check")
	var currentStatus ServiceStatus
	var httpStatusCode int
	error := db.Ping()
	if error == nil {
		currentStatus = UP
		httpStatusCode = http.StatusOK
	} else {
		log.Logger.Error("unable to connect to DB")
		currentStatus = DOWN
		httpStatusCode = http.StatusFailedDependency
	}
	status := &StatusResponse{
		Status:      currentStatus,
		ServiceName: h.ServiceName,
		UpTime:      h.UpTime,
		Environment: h.Environment,
		Version:     h.Version,
	}
	c.JSON(httpStatusCode, status)
}
