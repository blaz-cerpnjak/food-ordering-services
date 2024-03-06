package com.blazc.usermanagementapi

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.context.annotation.Bean
import org.yaml.snakeyaml.internal.Logger

@SpringBootApplication
class UserManagementApiApplication

val log = Logger.getLogger(UserManagementApiApplication::class.java.toString())

fun main(args: Array<String>) {
    runApplication<UserManagementApiApplication>(*args)
}

// Swagger http://localhost:8080/api/v1/swagger-ui/index.html
