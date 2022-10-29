package service

import (
	"encoding/json"

	"order-service/models"
	"order-service/pkg/kafka"
)

type PublisherService interface {
	Publish(order models.Order) error
}

type publisherImpl struct {
	producer kafka.Producer
	topic    string
}

func NewPublisher(p kafka.Producer, topic string) PublisherService {
	return &publisherImpl{
		producer: p,
		topic:    topic,
	}
}
func (p *publisherImpl) Publish(order models.Order) error {
	o, _ := json.Marshal(order)
	err := p.producer.Produce(nil, kafka.Message{
		Topic: p.topic,
		Key:   []byte(kafka.EventNameOrderCompleted),
		Value: o,
	})
	if err != nil {
		return err
	}
	return nil
}
