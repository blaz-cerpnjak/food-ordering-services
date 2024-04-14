package Converter

import (
	"API_GatewayWeb/DataStructures"
	pb "API_GatewayWeb/DataStructures/com.blazc"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func ConvertFromOrderStatusString(status string) (orderStatus pb.OrderStatus) {
	switch status {
	case "PENDING":
		orderStatus = pb.OrderStatus_PENDING
	case "PREPARING":
		orderStatus = pb.OrderStatus_PREPARING
	case "READY_FOR_PICKUP":
		orderStatus = pb.OrderStatus_READY_FOR_PICKUP
	case "DELIVERING":
		orderStatus = pb.OrderStatus_DELIVERING
	case "DELIVERED":
		orderStatus = pb.OrderStatus_DELIVERED
	case "CANCELLED":
		orderStatus = pb.OrderStatus_CANCELLED
	}
	return
}

func ConvertFromPaymentTypeString(paymentType string) (paymentTypeGrpc pb.PaymentType) {
	switch paymentType {
	case "CASH":
		paymentTypeGrpc = pb.PaymentType_CASH
	case "CREDIT_CARD":
		paymentTypeGrpc = pb.PaymentType_CREDIT_CARD
	}
	return

}

func ConvertOrderToGrpc(order DataStructures.Order) (orderGrpc *pb.Order) {

	itemsGrpc := make([]*pb.Product, 0)
	for _, item := range order.Items {
		itemsGrpc = append(itemsGrpc, &pb.Product{
			Id:           item.Id.Hex(),
			Name:         item.Name,
			Price:        item.Price,
			Image:        item.Image,
			RestaurantId: item.RestaurantId.Hex(),
		})
	}

	orderGrpc = &pb.Order{
		Id:               order.Id.Hex(),
		SellerId:         order.SellerId.Hex(),
		DeliveryPersonId: order.DeliveryPersonId.Hex(),
		Address:          order.Address,
		CustomerName:     order.CustomerName,
		Items:            itemsGrpc,
		Status:           ConvertFromOrderStatusString(order.Status),
		Timestamp:        time.Now().UTC().Unix(),
		PaymentType:      ConvertFromPaymentTypeString(order.PaymentType),
		TotalPrice:       order.TotalPrice,
	}

	return
}

func ConvertOrderFromGrpc(orderGrpc *pb.Order) (order DataStructures.Order, err error) {

	order.Id, err = primitive.ObjectIDFromHex(orderGrpc.Id)
	if err != nil {
		fmt.Println("orderId error: ", err.Error())
		return
	}

	order.SellerId, err = primitive.ObjectIDFromHex(orderGrpc.SellerId)
	if err != nil {
		fmt.Println("sellerId error: ", err.Error())
		return
	}

	order.DeliveryPersonId, err = primitive.ObjectIDFromHex(orderGrpc.DeliveryPersonId)
	if err != nil {
		fmt.Println("deliveryPersonId error: ", err.Error())
		return
	}

	order.Address = orderGrpc.Address
	order.CustomerName = orderGrpc.CustomerName

	items := make([]DataStructures.Product, 0)
	for _, item := range orderGrpc.Items {

		productId, err := primitive.ObjectIDFromHex(item.Id)
		if err != nil {
			fmt.Println("productId error: ", err.Error())
			return order, err
		}

		fmt.Println("restaurantId: ", item.RestaurantId)

		restaurantId, err := primitive.ObjectIDFromHex(item.RestaurantId)
		if err != nil {
			fmt.Println("restaurantId error: ", err.Error())
			return order, err
		}

		product := DataStructures.Product{
			Id:           productId,
			Name:         item.Name,
			Price:        item.Price,
			Image:        item.Image,
			RestaurantId: restaurantId,
		}

		items = append(items, product)
	}

	order.Items = items
	order.Status = orderGrpc.Status.String()
	order.Timestamp = orderGrpc.Timestamp
	order.PaymentType = orderGrpc.PaymentType.String()
	order.TotalPrice = orderGrpc.TotalPrice
	return
}
