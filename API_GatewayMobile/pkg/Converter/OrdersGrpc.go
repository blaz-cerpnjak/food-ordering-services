package Converter

import (
	"API_GatewayMobile/DataStructures"
	pb "API_GatewayMobile/DataStructures/com.blazc"
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

	orderItemsGrpc := make([]*pb.OrderItem, 0)
	for _, item := range order.OrderItems {
		orderItemsGrpc = append(orderItemsGrpc, &pb.OrderItem{
			Id:       item.Id.Hex(),
			Price:    item.Price,
			Quantity: item.Quantity,
			Product: &pb.Product{
				Id:    item.Id.Hex(),
				Name:  item.Product.Name,
				Price: item.Product.Price,
			},
		})
	}

	orderGrpc = &pb.Order{
		Id:               order.Id.Hex(),
		SellerId:         order.SellerId.Hex(),
		DeliveryPersonId: order.DeliveryPersonId.Hex(),
		Address:          order.Address,
		CustomerName:     order.CustomerName,
		Items:            orderItemsGrpc,
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
		return
	}

	order.SellerId, err = primitive.ObjectIDFromHex(orderGrpc.SellerId)
	if err != nil {
		return
	}

	order.DeliveryPersonId, err = primitive.ObjectIDFromHex(orderGrpc.DeliveryPersonId)
	if err != nil {
		return
	}

	order.Address = orderGrpc.Address
	order.CustomerName = orderGrpc.CustomerName

	items := make([]DataStructures.OrderItem, 0)
	for _, item := range orderGrpc.Items {
		id, err := primitive.ObjectIDFromHex(item.Id)
		if err != nil {
			return order, err
		}

		productId, err := primitive.ObjectIDFromHex(item.Product.Id)
		if err != nil {
			return order, err
		}

		orderItem := DataStructures.OrderItem{
			Id:       id,
			Price:    item.Price,
			Quantity: item.Quantity,
			Product: DataStructures.Product{
				Id:    productId,
				Name:  item.Product.Name,
				Price: item.Product.Price,
			},
		}
		items = append(items, orderItem)
	}

	order.OrderItems = items
	order.Status = orderGrpc.Status.String()
	order.Timestamp = orderGrpc.Timestamp
	order.PaymentType = orderGrpc.PaymentType.String()
	order.TotalPrice = orderGrpc.TotalPrice
	return
}
