package com.blazc.model

import io.quarkus.mongodb.panache.reactive.ReactivePanacheMongoEntity
import org.bson.types.ObjectId

class Product: ReactivePanacheMongoEntity() {
    var id: ObjectId? = null
    lateinit var restaurantId: ObjectId
    lateinit var name: String
    var price: Int = 0 // 100 = 1
}