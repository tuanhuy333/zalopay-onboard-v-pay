package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"order-service/mock"
	"order-service/pkg/kafka"
)

func TestNewMessageService(t *testing.T) {
	t.Run("new Message service success", func(t *testing.T) {
		s := NewMessageService(nil, nil)
		require.NotNil(t, s)
	})
}

func Test_messageImpl_HandleKafkaMessage(t *testing.T) {
	t.Run("handle kafka success", func(t *testing.T) {
		service := mock.NewMockOrderService(gomock.NewController(t))
		publishService := mock.NewMockPublisherService(gomock.NewController(t))

		c := messageImpl{
			service:          service,
			PublisherService: publishService,
		}

		service.EXPECT().UpdateOrderById(gomock.Any(), gomock.Any()).Return(nil, nil)
		publishService.EXPECT().Publish(gomock.Any()).Return(nil)

		err := c.HandleKafkaMessage(nil, kafka.Message{})
		require.NoError(t, err)
	})
	t.Run("handle kafka have update order error", func(t *testing.T) {
		service := mock.NewMockOrderService(gomock.NewController(t))
		publishService := mock.NewMockPublisherService(gomock.NewController(t))

		c := messageImpl{
			service:          service,
			PublisherService: publishService,
		}

		service.EXPECT().UpdateOrderById(gomock.Any(), gomock.Any()).Return(nil, errors.New("update order error"))
		//publishService.EXPECT().Publish(gomock.Any()).Return(nil)

		err := c.HandleKafkaMessage(nil, kafka.Message{})
		require.Error(t, err, "update order error")
	})
	t.Run("handle kafka have publish error", func(t *testing.T) {
		service := mock.NewMockOrderService(gomock.NewController(t))
		publishService := mock.NewMockPublisherService(gomock.NewController(t))

		c := messageImpl{
			service:          service,
			PublisherService: publishService,
		}

		service.EXPECT().UpdateOrderById(gomock.Any(), gomock.Any()).Return(nil, nil)
		publishService.EXPECT().Publish(gomock.Any()).Return(errors.New("publish error"))

		err := c.HandleKafkaMessage(nil, kafka.Message{})
		require.Error(t, err, "publish error")
	})
}
