package server

import (
	"strings"
	"time"

	"github.com/rameshsunkara/go-rest-api-example/internal/controllers"
	"github.com/rameshsunkara/go-rest-api-example/internal/models"
	"github.com/rameshsunkara/go-rest-api-example/pkg/log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(serviceInfo *models.ServiceMeta) *gin.Engine {
	if strings.Contains(serviceInfo.Environment, "dev") {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Middleware
	router := gin.New()
	router.Use(log.Ginzap(log.Logger, time.RFC3339, true))
	router.Use(log.RecoveryWithZap(log.Logger, true))

	// Routes

	// Routes - Health Check
	health := controllers.NewHealthController(serviceInfo)
	router.GET("/health", health.Status) // /health

	// Seed DB
	if serviceInfo.Environment == "dev" {
		seed := new(controllers.SeedDBController)
		seed.DBService.Prepare("ecommerce", "purchaseorders")
		router.POST("/seedDB", seed.SeedDB)
	}

	// Routes - API
	v1 := router.Group("/api/v1")
	{
		ordersGroup := v1.Group("orders")
		{
			orders := new(controllers.OrdersController)
			orders.DBService.Prepare("ecommerce", "purchaseorders")
			ordersGroup.GET("/", orders.GetAll)           // api/v1/orders
			ordersGroup.GET("/:id", orders.GetById)       // api/v1/orders/{id}
			ordersGroup.POST("/", orders.Post)            // api/v1/orders
			ordersGroup.PUT("/", orders.Post)             // api/v1/orders
			ordersGroup.DELETE("/:id", orders.DeleteById) // api/v1/orders/{id}
		}
	}

	// Routes - Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
