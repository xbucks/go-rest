package server

import (
	"go-rest-api-example/internal/controllers"
	"go-rest-api-example/pkg/log"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(environment string) *gin.Engine {
	gin.SetMode(gin.DebugMode)

	// Middleware
	router := gin.New()
	router.Use(log.Ginzap(log.Logger, time.RFC3339, true))
	router.Use(log.RecoveryWithZap(log.Logger, true))

	// Routes

	// Routes - Health Check
	health := new(controllers.HealthController)
	router.GET("/health", health.Status) // /health

	// Seed DB
	if environment == "dev" {
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
