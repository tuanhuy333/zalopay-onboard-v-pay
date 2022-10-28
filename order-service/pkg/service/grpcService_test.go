package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"order-service/mock"
	"order-service/pkg/order/pb"
)

func TestNewGRPCService(t *testing.T) {
	s := mock.NewMockOrderService(gomock.NewController(t))
	NewGRPCService(s)
}

func TestGRPCService_GetOrder(t *testing.T) {

	type fields struct {
		UnimplementedDisbursementServer pb.UnimplementedDisbursementServer
		orderService                    func(t *testing.T) OrderService
	}
	type args struct {
		ctx context.Context
		req *pb.GetOrderRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.Order
		wantErr bool
	}{
		{
			name: "get Order fail with validate error",
			args: args{req: nil},
			fields: fields{
				orderService: func(t *testing.T) OrderService {
					return nil
				},
			},
			wantErr: true,
		},
		{
			name: "get Order success",
			fields: fields{
				orderService: func(t *testing.T) OrderService {
					o := mock.NewMockOrderService(gomock.NewController(t))
					o.EXPECT().GetOrderById(gomock.Any()).Return(nil, nil)
					return o
				},
			},
			args:    args{req: &pb.GetOrderRequest{}},
			wantErr: false,
		},
		{
			name: "get Order error",
			fields: fields{
				orderService: func(t *testing.T) OrderService {
					o := mock.NewMockOrderService(gomock.NewController(t))
					o.EXPECT().GetOrderById(gomock.Any()).Return(nil, errors.New("error"))
					return o
				},
			},
			args:    args{req: &pb.GetOrderRequest{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &GRPCService{
				UnimplementedDisbursementServer: tt.fields.UnimplementedDisbursementServer,
				orderService:                    tt.fields.orderService(t),
			}
			got, err := s.GetOrder(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGRPCService_validate(t *testing.T) {
	type fields struct {
		UnimplementedDisbursementServer pb.UnimplementedDisbursementServer
		orderService                    OrderService
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &GRPCService{
				UnimplementedDisbursementServer: tt.fields.UnimplementedDisbursementServer,
				orderService:                    tt.fields.orderService,
			}
			if err := s.validate(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewGRPCService1(t *testing.T) {
	type args struct {
		o OrderService
	}
	tests := []struct {
		name string
		args args
		want *GRPCService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGRPCService(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGRPCService() = %v, want %v", got, tt.want)
			}
		})
	}
}
