package com.blazc.model

import org.bson.types.ObjectId

class Product {
    var id: ObjectId? = null
    lateinit var restaurantId: ObjectId
    lateinit var name: String
    lateinit var image: String
    var price: Int = 0 // 100 = 1
}