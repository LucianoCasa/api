package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucianocasa/api/internal/order"
)

func RegisterRoutes(r *gin.Engine, service order.Service) {
	r.GET("/order", func(c *gin.Context) {
		orders, err := service.ListOrders()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, orders)
	})

	r.POST("/order", func(c *gin.Context) {
		var o order.Order
		if err := c.ShouldBindJSON(&o); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := service.CreateOrder(&o); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, o)
	})
}
