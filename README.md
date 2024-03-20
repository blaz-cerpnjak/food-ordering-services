# Food Ordering System 📱 🚙 🍕

## Overview
This project serves as an example of food ordering system using a microservice approach.

## User Management API [REST]
Simple REST API for managing users, built using Spring Boot and Kotlin.

### How to run

```console
~$ cd UserManagementAPI
~$ docker-compose up
```

User Management API will be avaiable on ```localhost:8080/api/v1```. Access Swagger ```http://localhost:8080/api/v1/swagger-ui/index.html```.

### Unit Tests
Before running unit tests, ensure that you have a local MongoDB instance running on ```localhost:27017```.

### Endpoints
- GET /api/v1/health
- GET /api/v1/users
- GET /api/v1/users/{id}
- PUT /api/v1/users/{id}
- DELETE /api/v1/users/{id}

## Order Processing API [gRPC]
Simple gRPC API for managing users, built using Quarkus and Kotlin.

### How to run

```console
~$ cd OrderProcessingAPI
~$ docker-compose up
```

Order Processing gRPC API will be avaiable on ```grpc://localhost:9000```.

### Unit Tests
Before running unit tests, ensure that you have a local MongoDB instance running on ```localhost:27017```.

### Available gRPC Methods
```
service OrderService {
  rpc CreateOrder (Order) returns (Confirmation) {}
  rpc GetOrder (GetOrderRequest) returns (Order) {}
  rpc UpdateOrder (Order) returns (Confirmation) {}
  rpc GetOrdersBySeller (GetOrdersRequest) returns (stream Order) {}
  rpc GetOrdersByDeliveryPerson (GetOrdersRequest) returns (stream Order) {}
  rpc DeleteOrder (DeleteOrderRequest) returns (Confirmation) {}
  rpc Health (Empty) returns (Confirmation) {}
}
```

## Restaurant Management API [Reactive REST]
Reactive REST api with reactive mongo db

### How to run

```console
~$ cd RestaurantManagementAPI
~$ docker-compose up
```

Restaurant Management REST API will be avaiable on ```grpc://localhost:8080```.

### Unit Tests
Before running unit tests, ensure that you have a local MongoDB instance running on ```localhost:27017```.
