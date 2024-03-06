package com.blazc.usermanagementapi.dto

import org.bson.types.ObjectId

data class User (
        val id: ObjectId,
        val name: String,
        val lastname: String,
        val email: String,
        val password: String,
        val phone: String,
        val role: String // "buyer", "seller", "delivery"
)