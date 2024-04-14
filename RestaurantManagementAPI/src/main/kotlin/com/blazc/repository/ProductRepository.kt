package com.blazc.repository

import com.blazc.model.Product
import com.blazc.model.Restaurant
import io.quarkus.mongodb.panache.reactive.ReactivePanacheMongoRepository
import io.smallrye.mutiny.Uni
import jakarta.enterprise.context.ApplicationScoped
import org.bson.types.ObjectId

@ApplicationScoped
class ProductRepository: ReactivePanacheMongoRepository<Product> {

    fun add(product: Product): Uni<Product> {
        return persist(product)
    }

    fun getAllByRestaurantId(restaurantId: ObjectId): Uni<List<Product>> {
        return find("restaurantId", restaurantId).list()
    }

}