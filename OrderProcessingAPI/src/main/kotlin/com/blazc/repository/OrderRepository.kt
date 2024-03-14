package com.blazc.repository

import com.blazc.model.order.Order
import io.quarkus.mongodb.panache.PanacheMongoRepository
import jakarta.enterprise.context.ApplicationScoped
import org.bson.types.ObjectId

@ApplicationScoped
class OrderRepository : PanacheMongoRepository<Order> {

    fun findAllBySellerId(sellerId: ObjectId): List<Order> {
        return find("sellerId", sellerId).list()
    }

    fun findAllByDeliveryPersonId(deliveryPersonId: ObjectId): List<Order> {
        return find("deliveryPersonId", deliveryPersonId).list()
    }

}