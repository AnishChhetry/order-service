package handlers

import (
	"context"
	"log"
	"net/http"
	"order-service/models"
	"order-service/services"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	Service   services.OrderService
	OrderChan chan<- models.Order
}

func (h *OrderHandler) HealthStatus(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

func (h *OrderHandler) OrderList(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 500*time.Millisecond)
	defer cancel()

	orders, err := h.Service.GetOrderList(ctx)
	if err != nil {
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request took too long to respond"})
		return
	}
	c.JSON(http.StatusOK, orders)

}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var newOrder models.Order

	if err := c.BindJSON(&newOrder); err != nil {
		return
	}

	newOrder, err := h.Service.Create(c.Request.Context(), newOrder)
	if err != nil {
		switch err.Error() {
		case "Item with same ID already exists":
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		case "Item Name is Required":
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	}
	select {
	case h.OrderChan <- newOrder:
	default:
		log.Println("Warning: Worker is busy, dropping notigication for order: ", newOrder.ID)
	}

	c.JSON(http.StatusCreated, newOrder)
}
