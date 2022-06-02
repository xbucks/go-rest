package server

import (
	"sync"

	"github.com/rameshsunkara/go-rest-api-example/internal/config"
	"github.com/rameshsunkara/go-rest-api-example/internal/models"
)

var runOnce sync.Once

func Init(serviceInfo *models.ServiceMeta) {
	config := config.GetConfig()
	port := config.GetString("server.port")
	runOnce.Do(func() {
		startServer(port, serviceInfo)
	})

}

func startServer(port string, serviceInfo *models.ServiceMeta) {
	r := NewRouter(serviceInfo)
	r.Run(":" + port)
}
