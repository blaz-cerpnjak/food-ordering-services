package com.blazc.usermanagementapi

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
class UserManagementApiApplication

fun main(args: Array<String>) {
    runApplication<UserManagementApiApplication>(*args)
}

// Swagger http://localhost:8080/api/v1/swagger-ui/index.html
