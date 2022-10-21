package service

import (
	"context"
	"encoding/json"
	"log"

	"V_Pay_Onboard_Program/models"
	"V_Pay_Onboard_Program/pkg/kafka"
)

type HandleMessageService interface {
	HandleKafkaMessage(ctx context.Context, msg kafka.Message) error
}

type messageImpl struct {
	service          *Storage
	PublisherService PublisherService
}

func NewMessageService(s *Storage, p PublisherService) HandleMessageService {
	return &messageImpl{service: s, PublisherService: p}
}

func (s *messageImpl) HandleKafkaMessage(ctx context.Context, msg kafka.Message) error {
	var o models.Order
	json.Unmarshal(msg.Value, &o)
	_, err := s.service.UpdateOrderById(o.OrderNo, &o)
	if err != nil {
		return err
	}
	log.Printf("Updated Order: %v", &o)
	// publish event ordercompleted
	if err := s.PublisherService.Publish(o); err != nil {
		return err
	}
	log.Printf("Publish message in orders topic")
	return nil
}
