package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rameshsunkara/go-rest-api-example/internal/db"
	"github.com/rameshsunkara/go-rest-api-example/pkg/log"
)

type HealthController struct{}

// Status godoc
// @Summary      Status/Health of the service (Doesn't work in Swagger UI)
// @Description  Provides the current status/health of the service. It checks the health of all the critical components of the service.
// @Tags         Status
// @Accept       json
// @Produce      plain
// @Success      200  {string}  string  "OK"
// @Failure      503  {string}  string  "Something is wrong"
// @Router       /health [get]
func (h *HealthController) Status(c *gin.Context) {
	log.Logger.Debug("In Status Check")
	err := db.Ping()
	if err != nil {
		c.String(http.StatusServiceUnavailable, "Something is wrong")
		return
	}
	c.String(http.StatusOK, "OK")
	return
}
