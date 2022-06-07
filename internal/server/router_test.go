package server

import (
	"context"
	"github.com/rameshsunkara/go-rest-api-example/internal/models"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	svcInfo = &models.ServiceInfo{
		Name:        "test-api-service",
		Version:     "rams-fav",
		UpTime:      time.Now(),
		Environment: "test",
	}
)

type MockMongoDBClient struct{}

func (m *MockMongoDBClient) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return nil
}

type MockMongoDataBase struct{}

func (m *MockMongoDataBase) Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection {
	return nil
}

func TestListOfRoutes(t *testing.T) {
	router := WebRouter(svcInfo, &MockMongoDBClient{}, &MockMongoDataBase{})
	list := router.Routes()

	assertRoutePresent(t, list, gin.RouteInfo{
		Method: http.MethodGet,
		Path:   "/status",
	})

	assertRoutePresent(t, list, gin.RouteInfo{
		Method: http.MethodGet,
		Path:   "/api/v1/orders",
	})

	assertRoutePresent(t, list, gin.RouteInfo{
		Method: http.MethodGet,
		Path:   "/api/v1/orders/:id",
	})

	assertRoutePresent(t, list, gin.RouteInfo{
		Method: http.MethodPost,
		Path:   "/api/v1/orders",
	})

	assertRoutePresent(t, list, gin.RouteInfo{
		Method: http.MethodPut,
		Path:   "/api/v1/orders",
	})

	assertRoutePresent(t, list, gin.RouteInfo{
		Method: http.MethodDelete,
		Path:   "/api/v1/orders/:id",
	})

}

func assertRoutePresent(t *testing.T, gotRoutes gin.RoutesInfo, wantRoute gin.RouteInfo) {
	for _, gotRoute := range gotRoutes {
		if gotRoute.Path == wantRoute.Path && gotRoute.Method == wantRoute.Method {
			return
		}
	}
	t.Errorf("route not found: %v", wantRoute)
}
