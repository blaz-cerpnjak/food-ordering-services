package DataStructures

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Lastname string             `json:"lastname" bson:"lastname"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	Phone    string             `json:"phone" bson:"phone"`
	Role     string             `json:"role" bson:"role"`
}

type Restaurant struct {
	Id      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name    string             `json:"name" bson:"name"`
	Address string             `json:"address" bson:"address"`
	Contact string             `json:"contact" bson:"contact"`
}

type Product struct {
	Id           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	RestaurantId primitive.ObjectID `json:"restaurantId" bson:"restaurantId"`
	Name         string             `json:"name" bson:"name"`
	Price        int32              `json:"price" bson:"price"` // in cents
}

type Order struct {
	Id               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SellerId         primitive.ObjectID `json:"sellerId" bson:"sellerId"`
	DeliveryPersonId primitive.ObjectID `json:"deliveryPersonId" bson:"deliveryPersonId"`
	Address          string             `json:"address" bson:"address"`
	CustomerName     string             `json:"customerName" bson:"customerName"`
	OrderItems       []OrderItem        `json:"items" bson:"items"`
	Status           string             `json:"status" bson:"status"`
	Timestamp        int64              `json:"timestamp" bson:"timestamp"`
	PaymentType      string             `json:"paymentType" bson:"paymentType"`
	TotalPrice       int32              `json:"totalPrice" bson:"totalPrice"` // in cents
}

type OrderItem struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Product  Product            `json:"product" bson:"product"`
	Quantity int32              `json:"quantity" bson:"quantity"`
	Price    int32              `json:"price" bson:"price"` // in cents
}
