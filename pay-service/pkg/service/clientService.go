package service

import (
	"context"
	"fmt"

	"pay-service/pkg/disbursement/pb"
)

type ClientInterface interface {
	CallGetOrder(context context.Context, orderId int) (*pb.Order, error)
}
type clientGrpc struct {
	client pb.DisbursementClient
}

func NewClient(client pb.DisbursementClient) ClientInterface {
	return &clientGrpc{
		client: client,
	}
}
func (c *clientGrpc) CallGetOrder(context context.Context, orderId int) (*pb.Order, error) {
	fmt.Println("Call get Order")
	request := pb.GetOrderRequest{OrderNo: int64(orderId)}

	o, err := c.client.GetOrder(context, &request)
	if err != nil {
		return nil, err
	}
	return o, nil
}
