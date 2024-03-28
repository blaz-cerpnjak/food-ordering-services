package Logic

import (
	"API_GatewayWeb/DataStructures"
	pb "API_GatewayWeb/DataStructures/com.blazc"
	"API_GatewayWeb/pkg/Converter"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *Controller) Health(ctx context.Context) (status string, err error) {
	confirmation, err := c.grpc.Client.Health(ctx, &pb.Empty{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	status = confirmation.Message
	return
}

func (c *Controller) CreateOrder(ctx context.Context, order DataStructures.Order) (confirmation *pb.Confirmation, err error) {

	orderGrpc := Converter.ConvertOrderToGrpc(order)

	confirmation, err = c.grpc.Client.CreateOrder(ctx, orderGrpc)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}

func (c *Controller) UpdateOrder(ctx context.Context, id primitive.ObjectID) (confirmation *pb.Confirmation, err error) {
	confirmation, err = c.grpc.Client.UpdateOrder(ctx, &pb.Order{Id: id.Hex()})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}

func (c *Controller) DeleteOrder(ctx context.Context, id primitive.ObjectID) (confirmation *pb.Confirmation, err error) {
	confirmation, err = c.grpc.Client.DeleteOrder(ctx, &pb.DeleteOrderRequest{Id: id.Hex()})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return
}

func (c *Controller) GetOrdersBySellerId(ctx context.Context, id primitive.ObjectID) (orders []DataStructures.Order, err error) {
	orderStream, err := c.grpc.Client.GetOrdersBySeller(ctx, &pb.GetOrdersRequest{Id: id.Hex()})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	orders = make([]DataStructures.Order, 0)

	for {
		orderGrpc, err := orderStream.Recv()
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		order, err := Converter.ConvertOrderFromGrpc(orderGrpc)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		orders = append(orders, order)
	}

	return
}
