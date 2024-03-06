package com.blazc.usermanagementapi.dto.post

data class PostUser (
    val name: String,
    val lastname: String,
    val email: String,
    val password: String,
    val phone: String,
    val role: String // "buyer", "seller", "delivery"
)