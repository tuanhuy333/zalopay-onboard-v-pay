package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"order-service/mock"
	"order-service/models"
)

func TestNewPublisher(t *testing.T) {

	t.Run("new publisher success", func(t *testing.T) {
		p := NewPublisher(nil, "")
		require.NotNil(t, p, "not nil")
	})
}

func Test_publisherImpl_Publish(t *testing.T) {
	t.Run("publisher success", func(t *testing.T) {
		p := mock.NewMockProducer(gomock.NewController(t))
		s := publisherImpl{producer: p}

		m := models.Order{}
		p.EXPECT().Produce(gomock.Any(), gomock.Any()).Return(nil)

		s.Publish(m)
	})
	t.Run("publisher failed", func(t *testing.T) {
		p := mock.NewMockProducer(gomock.NewController(t))
		s := publisherImpl{producer: p}

		m := models.Order{}
		p.EXPECT().Produce(gomock.Any(), gomock.Any()).Return(errors.New("producer fail"))

		err := s.Publish(m)
		require.Error(t, err, "producer fail")
	})

}
