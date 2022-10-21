package kafka

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

var (
	ErrConsumerGroupClosed = sarama.ErrClosedConsumerGroup
	ErrNoTopicProvided     = errors.New("no topic provided")
)

// Consumer represents a Sarama consumer group consumer
type Consumer interface {
	Consume(topic string, fn HandlerFunc)
	Start() error
	Stop() error
}

type consumer struct {
	cg      sarama.ConsumerGroup
	handler *handler
}

type handler struct {
	mux    sync.Mutex
	topics map[string]HandlerFunc
}

type HandlerFunc func(ctx context.Context, msg Message) error

func NewConsumer(brokers []string, groupID string) (Consumer, error) {
	cg, err := sarama.NewConsumerGroup(brokers, groupID, sarama.NewConfig())
	if err != nil {
		return nil, fmt.Errorf("kafka: create group consumer: %v", err)
	}

	return &consumer{
		cg: cg,
		handler: &handler{
			topics: make(map[string]HandlerFunc),
		},
	}, nil
}

func (c *consumer) Consume(topic string, fn HandlerFunc) {
	c.handler.mux.Lock()
	c.handler.topics[topic] = fn
	c.handler.mux.Unlock()
}

func (c *consumer) Start() error {
	topics := make([]string, 0, len(c.handler.topics))
	for t, h := range c.handler.topics {
		if h == nil {
			return fmt.Errorf("handler of topic: %s is nil", t)
		}

		topics = append(topics, t)
	}

	if len(topics) == 0 {
		return ErrNoTopicProvided
	}

	for {
		err := c.cg.Consume(context.TODO(), topics, c.handler)
		if errors.Is(err, sarama.ErrClosedConsumerGroup) {
			return ErrConsumerGroupClosed
		}

		if err != nil {
			//logging.FromContext(context.TODO()).Errorf("kafka: consumer group error: %v", err)
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func (c *consumer) Stop() error {
	return c.cg.Close()
}

func (h *handler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *handler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (h *handler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/main/consumer_group.go#L27-L29
	for {
		select {
		case message := <-claim.Messages():
			h.consume(session.Context(), message)
			session.MarkMessage(message, "")

		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/Shopify/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}

func (h *handler) consume(ctx context.Context, msg *sarama.ConsumerMessage) {
	f, ok := h.topics[msg.Topic]
	if !ok {
		return
	}

	headers := make(map[string]string)
	for _, h := range msg.Headers {
		if h == nil {
			continue
		}

		headers[string(h.Key)] = string(h.Value)
	}

	_ = f(ctx, Message{
		Topic:   msg.Topic,
		Key:     msg.Key,
		Value:   msg.Value,
		Headers: headers,
	})
}
