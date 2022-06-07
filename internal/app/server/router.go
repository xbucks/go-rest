package server

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/rameshsunkara/go-rest-api-example/internal/db"
	"github.com/rameshsunkara/go-rest-api-example/internal/handlers"
	"github.com/rameshsunkara/go-rest-api-example/internal/models"
	"github.com/rameshsunkara/go-rest-api-example/pkg/util"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func WebRouter(svcInfo *models.ServiceInfo, client handlers.MongoDBClient, db db.MongoDBDatabase) (router *gin.Engine) {
	ginMode := gin.ReleaseMode
	if util.IsDevMode(svcInfo.Environment) {
		ginMode = gin.DebugMode
		gin.ForceConsoleColor()
	}
	gin.SetMode(ginMode)

	// Middleware
	router = gin.Default()
	pprof.Register(router) // TODO: Add debug routes only for Admins /debug/*
	// TODO: Enforce there is authorization information with applicable requests
	// TODO: log everything from gin in json

	// Routes

	// Routes - Status Check
	status := handlers.NewStatusHandler(svcInfo, client)
	router.GET("/status", status.CheckStatus) // /status

	// Seed DB
	if util.IsDevMode(svcInfo.Environment) {
		seed := new(handlers.SeedDBController)
		router.POST("/seedDB", seed.SeedDB)
	}

	// Routes - API
	v1 := router.Group("/api/v1")
	{
		ordersGroup := v1.Group("orders")
		{
			orders := handlers.NewOrdersHandler(db)
			ordersGroup.GET("/", orders.GetAll)           // api/v1/orders
			ordersGroup.GET("/:id", orders.GetById)       // api/v1/orders/{id}
			ordersGroup.POST("/", orders.Post)            // api/v1/orders
			ordersGroup.PUT("/", orders.Post)             // api/v1/orders
			ordersGroup.DELETE("/:id", orders.DeleteById) // api/v1/orders/{id}
		}
	}

	// Routes - Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return
}
