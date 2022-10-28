package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"pay-service/models"
	"pay-service/pkg/service"
)

type Handler struct {
	Service          *service.Storage
	PublisherService service.PublisherService
	Client           service.ClientInterface
}

func (h *Handler) GetOrderById(context *gin.Context) {

	id := context.Query("orderNo")
	orderId, errParse := strconv.Atoi(id)
	if errParse != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "orderNo is invalid"})
		context.Abort()
		return
	}
	o, err := h.Client.CallGetOrder(context, orderId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": o})

}
func (h *Handler) Confirm(context *gin.Context) {
	var request models.Order
	// validate
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	log.Printf("request %v", request)
	if err := h.PublisherService.Publish(request); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"data": "Sended Event"})

}
