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
}

func NewPublisher(p kafka.Producer) PublisherService {
	return &publisherImpl{
		producer: p,
	}
}

func (p *publisherImpl) Publish(order models.Order) error {
	o, err := json.Marshal(order)
	if err != nil {
		return err
	}
	p.producer.Produce(nil, kafka.Message{
		Topic: "local.v.pay.payment",
		Key:   []byte(kafka.EventNamePaymentFinished),
		Value: o,
	})
	return nil
}
