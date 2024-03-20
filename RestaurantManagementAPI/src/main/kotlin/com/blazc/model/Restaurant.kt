package com.blazc.model

import org.bson.types.ObjectId

class Restaurant {
    var id: ObjectId? = null
    lateinit var name: String
    lateinit var address: String
    lateinit var contact: String
}
