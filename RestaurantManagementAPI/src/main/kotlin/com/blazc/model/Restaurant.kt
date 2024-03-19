package com.blazc.model

import io.quarkus.mongodb.panache.reactive.ReactivePanacheMongoEntity
import org.bson.types.ObjectId

class Restaurant: ReactivePanacheMongoEntity() {
    var id: ObjectId? = null
    lateinit var name: String
    lateinit var address: String
    lateinit var contact: String
}
