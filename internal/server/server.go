package server

import (
	"github.com/rameshsunkara/go-rest-api-example/internal/db"
	"github.com/rameshsunkara/go-rest-api-example/internal/handlers"
	"sync"

	"github.com/rameshsunkara/go-rest-api-example/internal/config"
	"github.com/rameshsunkara/go-rest-api-example/internal/models"
)

var runOnce sync.Once

func Init(serviceInfo *models.ServiceInfo, client handlers.MongoDBClient, db db.MongoDBDatabase) {
	config := config.GetConfig()
	port := config.GetString("server.port")
	runOnce.Do(func() {
		startServer(port, serviceInfo, client, db)
	})

}

func startServer(port string, serviceInfo *models.ServiceInfo, client handlers.MongoDBClient, db db.MongoDBDatabase) {
	r := WebRouter(serviceInfo, client, db)
	r.Run(":" + port)
}
