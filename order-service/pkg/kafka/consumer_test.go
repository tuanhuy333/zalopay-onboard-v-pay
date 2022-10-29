package kafka

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewConsumer(t *testing.T) {
	t.Run("connect to wrong brokers", func(t *testing.T) {
		c, err := NewConsumer([]string{""}, "")
		require.Error(t, err, "should return err not nil")
		require.Nil(t, c)
	})
}

func TestConsume(t *testing.T) {
	t.Run("consume with handle func return nil error", func(t *testing.T) {
		c := consumer{
			handler: &handler{
				topics: make(map[string]HandlerFunc),
			},
		}
		c.Consume("success-topic", func(ctx context.Context, msg Message) error {
			return nil
		})
		require.NoError(t, c.handler.topics["success-topic"](nil, Message{}))
	})

	t.Run("consume with handle func return an error", func(t *testing.T) {
		c := consumer{
			handler: &handler{
				topics: make(map[string]HandlerFunc),
			},
		}
		c.Consume("failed-topic", func(ctx context.Context, msg Message) error {
			return errors.New("consume msg failed")
		})
		require.Error(t, c.handler.topics["failed-topic"](nil, Message{}))
	})
}

// Currently, sarama does not have a mock consumer group, so it's hard to write a unit test.
// Related PR: https://github.com/Shopify/sarama/pull/1750
func TestStartConsumer(t *testing.T) {
	t.Run("start a consumer dont has any topic should return ErrNoTopicProvided", func(t *testing.T) {
		c := consumer{
			handler: &handler{
				topics: nil,
			},
		}
		err := c.Start()
		require.Equal(t, err, ErrNoTopicProvided)
	})

	t.Run("start a consumer has a topic without handle func should return err", func(t *testing.T) {
		c := consumer{
			handler: &handler{
				topics: map[string]HandlerFunc{},
			},
		}
		c.Consume("test-topic", nil)
		err := c.Start()
		require.Equal(t, err.Error(), "handler of topic: test-topic is nil")
	})
}
