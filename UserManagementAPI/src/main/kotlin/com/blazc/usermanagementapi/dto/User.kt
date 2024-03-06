package com.blazc.usermanagementapi.dto

data class User (
        val id: String,
        val name: String,
        val lastname: String,
        val email: String,
        val password: String,
        val phone: String,
        val role: String // "buyer", "seller", "delivery"
)