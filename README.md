# Food Ordering System

### TODO:
- RestaurantsAPI
- OrdersProcessingAPI

## Overview
This project serves as an example of food ordering system using a microservice approach.

## UserManagementAPI
Simple microservice for managing users, built using Spring Boot and Kotlin.

### How to run

```console
~$ docker-compose up
```

UserManagementAPI will be avaiable on ```localhost:8080/api/v1```. Access Swagger ```http://localhost:8080/api/v1/swagger-ui/index.html```.

### Unit Tests
Before running unit tests, ensure that you have a local MongoDB instance running on ```localhost:27017```.

### Endpoints
GET /api/v1/health
GET /api/v1/users
GET /api/v1/users/{id}
PUT /api/v1/users/{id}
DELETE /api/v1/users/{id}
