package kafka

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
)

type Producer interface {
	Produce(_ context.Context, msg Message) error
	Close() error
}

type producer struct {
	p sarama.SyncProducer
}

func NewProducer(brokers []string) (Producer, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true

	p, err := sarama.NewSyncProducer(brokers, cfg)
	if err != nil {
		return nil, fmt.Errorf("kafka: create sync producer (brokers=%v):%v", brokers, err)
	}
	return &producer{p: p}, nil
}

func (p *producer) Produce(_ context.Context, msg Message) error {
	if msg.Headers == nil {
		msg.Headers = make(map[string]string)
	}
	headers := make([]sarama.RecordHeader, 0, len(msg.Headers))
	for k, v := range msg.Headers {
		headers = append(headers, sarama.RecordHeader{
			Key:   []byte(k),
			Value: []byte(v),
		})

	}
	_, _, err := p.p.SendMessage(&sarama.ProducerMessage{
		Topic:    msg.Topic,
		Key:      sarama.ByteEncoder(msg.Key),
		Value:    sarama.ByteEncoder(msg.Value),
		Headers:  headers,
		Metadata: nil,
	})
	if err != nil {
		return fmt.Errorf("kafka: produce message (topic=%v): %v", msg.Topic, err)
	}
	return nil
}

func (p *producer) Close() error {
	return p.p.Close()
}
