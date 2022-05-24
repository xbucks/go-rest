package controllers

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-example/internal/db"
	"go-rest-api-example/internal/models"
	"go-rest-api-example/pkg/log"
	"net/http"
)

type OrdersController struct {
	DBService db.OrdersCrudService
}

// Post  godoc
// @Summary      Creates or Updates an order
// @Description  Used to either create or update an order
// @Tags         Fetch
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /orders/ [post]
func (ordersController *OrdersController) Post(c *gin.Context) {
	purchaseRequest := models.Order{}

	if err := c.BindJSON(&purchaseRequest); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if purchaseRequest.ID.IsZero() {
		if uid, _ := ordersController.DBService.CreateOrder(&purchaseRequest); uid != nil {
			c.JSON(http.StatusOK, uid)
			return
		}
	} else {
		if updatedCount, _ := ordersController.DBService.UpdateOrder(&purchaseRequest); updatedCount != 0 {
			c.JSON(http.StatusOK, updatedCount)
			return
		}
	}

	c.JSON(http.StatusInternalServerError, "Unexpected Error occurred")
}

// GetAll  godoc
// @Summary      Fetch all orders
// @Description  Fetches all orders
// @Tags         Fetch
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /orders/ [get]
func (ordersController *OrdersController) GetAll(c *gin.Context) {
	log.Logger.Debug("fetch all documents of purchase orders")
	orders, err := ordersController.DBService.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error occurred while retrieved purchase orders", "error": err})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders": orders})
	return
}

// GetById  godoc
// @Summary      Fetch single Order document identified by give id
// @Description  Fetch single Order document identified by give id
// @Param        id   path      string  true  "Order ID"
// @Tags         Fetch
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500            {string}  string  "bad request"
// @Router       /orders/{id} [get]
func (ordersController *OrdersController) GetById(c *gin.Context) {
	if c.Param("id") != "" {
		order, err := ordersController.DBService.GetOrderByID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve order details", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"Order": order})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}

// DeleteById  godoc
// @Summary      Delete single Order document identified by give id
// @Description  Delete single Order document identified by give id
// @Param        id   path      string  true  "Order ID"
// @Tags         Fetch
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500            {string}  string  "bad request"
// @Router       /orders/{id} [delete]
func (ordersController *OrdersController) DeleteById(c *gin.Context) {
	if c.Param("id") != "" {
		count, err := ordersController.DBService.DeleteOrderByID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to retrieve order details", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"DeletedCount": count})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
	return
}
