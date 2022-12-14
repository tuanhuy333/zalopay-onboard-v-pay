package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"order-service/models"
	"order-service/pkg/service"
	"order-service/util/authutil"
)

type Handler struct {
	Service          service.OrderService
	PublisherService service.PublisherService
}

func (h *Handler) CreateOrders(context *gin.Context) {
	var request models.Order

	// validate
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// validate mac
	if !authutil.ValidMAC("key", request.Mac, request.AppID) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "mac not valid"})
		context.Abort()
		return
	}

	// implement
	err := h.Service.CreateOrder(&request)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// publish message
	//h.PublisherService.Publish()

	// show success
	url := fmt.Sprintf("http://localhost:8098/order?orderNo=%v", request.OrderNo)
	context.JSON(http.StatusOK, gin.H{"url": url})
}

func (h *Handler) GetOrders(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "http://localhost:3001")

	var orders []models.Order
	o, err := h.Service.GetAllOrder(&orders)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": o})

}
