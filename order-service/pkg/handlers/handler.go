package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

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
	//var user models.User
	// validate
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// validate mac
	if !authutil.ValidMAC("key", "8bbc27fa3bd74d7c55f7eda2400213ce30b3434b54909557dc7115aa8f454214", "a", "b") {
		context.JSON(http.StatusBadRequest, gin.H{"error": errors.New("mac not valid")})
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
	url := fmt.Sprintf("http://localhost:8098/order/%v", request.OrderNo)
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

func (h *Handler) GetOrderById(context *gin.Context) {
	context.Header("Access-Control-Allow-Origin", "http://localhost:8098")

	id := context.Param("id")
	orderId, _ := strconv.Atoi(id)
	o, err := h.Service.GetOrderById(orderId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": o})

}
