package service

import (
	"context"

	"order-service/pkg/disbursement/converters"
	"order-service/pkg/disbursement/pb"
)

type GRPCService struct {
	pb.UnimplementedDisbursementServer
	orderService *Storage
}

func NewGRPCService(o *Storage) *GRPCService {
	return &GRPCService{
		orderService: o,
	}
}

//func (s *GRPCService) validateCreateOrderReq(req *pb.CreateOrderRequest) error {
//	if err := s.validate(req); err != nil {
//		return err
//	}
//
//	return nil
//}

func (s *GRPCService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.Order, error) {
	//if err := s.validate(req); err != nil {
	//	return nil, err
	//}

	orderCore, err := s.orderService.GetOrderById(int(req.OrderNo))
	if err != nil {
		return nil, err
	}

	return converters.OrderCore2Pb(orderCore)
}

//
//func (s *GRPCService) validate(v interface{}) error {
//	if reflect.ValueOf(v).IsNil() {
//		return errors.New(errors.CodeInvalidArgument, "invalid request:nil")
//	}
//
//	err := s.validator.Struct(v)
//	if err != nil {
//		var ve validator.ValidationErrors
//		if stderr.As(err, &ve) {
//			err = errors.Newf(errors.CodeInvalidArgument, "Invalid field(s): %v", validatorutil.HandleValidationError(ve))
//		}
//	}
//
//	return err
//}
