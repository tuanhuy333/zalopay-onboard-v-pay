package kafka

import (
	"context"
	"errors"
	"testing"

	"github.com/Shopify/sarama"
	"github.com/Shopify/sarama/mocks"
	"github.com/stretchr/testify/require"
)

func Test_producer_Produce(t *testing.T) {
	sendErr := errors.New("send failed")
	type fields struct {
		p sarama.SyncProducer
	}
	type args struct {
		ctx context.Context
		msg Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "send msg success",
			fields: fields{
				p: mocks.NewSyncProducer(t, nil).ExpectSendMessageAndSucceed(),
			},
			args: args{
				msg: Message{Topic: "test", Key: []byte("123"), Value: []byte("test")},
			},
			wantErr: false,
		},
		{
			name: "send msg with header success",
			fields: fields{
				p: mocks.NewSyncProducer(t, nil).ExpectSendMessageAndSucceed(),
			},
			args: args{
				msg: Message{Topic: "test", Key: []byte("123"), Value: []byte("test"), Headers: map[string]string{"foo": "bar"}},
			},
			wantErr: false,
		},
		{
			name: "send msg failed",
			fields: fields{
				p: mocks.NewSyncProducer(t, nil).ExpectSendMessageAndFail(sendErr),
			},
			args: args{
				msg: Message{Topic: "test", Key: []byte("123"), Value: []byte("test")},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &producer{
				p: tt.fields.p,
			}
			if err := p.Produce(tt.args.ctx, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Produce() error = %v, wantErr %v", err, tt.wantErr)
			}
			require.NoError(t, p.Close())
		})
	}
}

func TestNewProducer(t *testing.T) {
	t.Run("connect to wrong brokers", func(t *testing.T) {
		c, err := NewProducer([]string{""})
		require.Error(t, err, "should return err not nil")
		require.Nil(t, c)
	})

}
