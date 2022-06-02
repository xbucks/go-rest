package server

import (
	"sync"

	"github.com/rameshsunkara/go-rest-api-example/internal/config"
)

var runOnce sync.Once

func Init(environment string) {
	config := config.GetConfig()
	port := config.GetString("server.port")
	runOnce.Do(func() {
		startServer(port, environment)
	})

}

func startServer(port string, environment string) {
	r := NewRouter(environment)
	r.Run(":" + port)
}
