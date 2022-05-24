package server

import (
	"go-rest-api-example/internal/config"
	"sync"
)

var runOnce sync.Once

func Init() {
	config := config.GetConfig()
	port := config.GetString("server.port")
	runOnce.Do(func() {
		startServer(port)
	})

}

func startServer(port string) {
	r := NewRouter()
	r.Run(":" + port)
}
