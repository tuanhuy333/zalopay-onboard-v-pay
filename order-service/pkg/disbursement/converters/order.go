package converters

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"order-service/models"
	"order-service/pkg/disbursement/pb"
)

// OrderCore2Pb convert Order core to protobuf message
func OrderCore2Pb(order *models.Order) (*pb.Order, error) {
	if order == nil {
		return nil, nil
	}

	return &pb.Order{
		OrderNo:    int32(order.OrderNo),
		MerchantId: order.MerchantID,
		AppId:      order.AppID,

		Description: order.Description,
		Amount:      float32(order.Amount),
		ProductCode: order.ProductCode,

		Status:     order.Status,
		CreateTime: timestamppb.New(order.CreateTime),
	}, nil
}
