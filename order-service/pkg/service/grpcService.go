package service

import (
	"context"
	"errors"
	"reflect"

	"order-service/pkg/order/converters"
	"order-service/pkg/order/pb"
)

type GRPCService struct {
	pb.UnimplementedDisbursementServer
	orderService OrderService
}

func NewGRPCService(o OrderService) *GRPCService {
	return &GRPCService{
		orderService: o,
	}
}

func (s *GRPCService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.Order, error) {
	if err := s.validate(req); err != nil {
		return nil, err
	}

	orderCore, err := s.orderService.GetOrderById(int(req.OrderNo))
	if err != nil {
		return nil, err
	}

	return converters.OrderCore2Pb(orderCore)
}

func (s *GRPCService) validate(v interface{}) error {
	if reflect.ValueOf(v).IsNil() {
		return errors.New("invalid request:nil")
	}

	return nil
}
