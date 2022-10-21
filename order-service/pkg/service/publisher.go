package service

import (
	"encoding/json"

	"V_Pay_Onboard_Program/models"
	"V_Pay_Onboard_Program/pkg/kafka"
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
	o, err := json.Marshal(order)
	if err != nil {
		return err
	}
	p.producer.Produce(nil, kafka.Message{
		Topic: p.topic,
		Key:   []byte(kafka.EventNameOrderCompleted),
		Value: o,
	})
	return nil
}
