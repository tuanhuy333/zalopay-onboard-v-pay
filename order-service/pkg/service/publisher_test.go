package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"order-service/mock"
	"order-service/models"
	"order-service/pkg/kafka"
)

type producerServiceMock struct {
}

//
//func (p *producerServiceMock) Produce(context context.Context, msg kafka.Message) error {
//	return nil
//}
//func (p *producerServiceMock) Close() error {
//	return nil
//}

func TestNewPublisher(t *testing.T) {

	t.Run("new publisher success", func(t *testing.T) {
		p := NewPublisher(nil, "")
		require.NotNil(t, p, "not nil")
	})
}

func Test_publisherImpl_Publish(t *testing.T) {

	type fields struct {
		producer kafka.Producer
		topic    string
	}
	type args struct {
		order models.Order
	}
	tests := []struct {
		name    string
		fields  func(t *testing.T) fields
		args    args
		wantErr bool
	}{

		{

			name: "publish success",
			fields: func(t *testing.T) fields {
				p := mock.NewMockProducer(gomock.NewController(t))
				p.EXPECT().Produce(gomock.Any(), gomock.Any()).Return(nil)

				return fields{producer: p}
				//producer:
				//	func() kafka.Producer {
				//		return nil
				//	}

			},

			args:    args{order: models.Order{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &publisherImpl{
				producer: tt.fields(t).producer,
				topic:    tt.fields(t).topic,
			}
			if err := p.Publish(tt.args.order); (err != nil) != tt.wantErr {
				t.Errorf("Publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func Test_publisherImpl_Publish1(t *testing.T) {
	p := mock.NewMockProducer(gomock.NewController(t))
	s := publisherImpl{producer: p}

	m := models.Order{}
	p.EXPECT().Produce(gomock.Any(), gomock.Any()).Return(nil)

	s.Publish(m)
}
