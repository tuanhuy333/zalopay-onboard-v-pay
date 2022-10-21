package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"V_Pay_Onboard_Program/models"
	"V_Pay_Onboard_Program/pkg/service"
)

type Handler struct {
	Service          *service.Storage
	PublisherService service.PublisherService
}

func (h *Handler) GetOrderById(context *gin.Context) {
	id := context.Param("id")
	o, err := h.Service.GetOrderById(id)
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
