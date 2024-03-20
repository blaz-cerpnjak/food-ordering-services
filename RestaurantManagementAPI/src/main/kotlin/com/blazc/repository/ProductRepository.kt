package com.blazc.repository

import com.blazc.model.Product
import io.quarkus.mongodb.panache.reactive.ReactivePanacheMongoRepository
import io.smallrye.mutiny.Uni
import jakarta.enterprise.context.ApplicationScoped
import org.bson.types.ObjectId

@ApplicationScoped
class ProductRepository: ReactivePanacheMongoRepository<Product> {

    fun create(product: Product): Uni<Product> {
        return persist(product).onItem().transform { product }
    }

    fun getAllByRestaurantId(restaurantId: ObjectId): Uni<List<Product>> {
        return find("restaurantId", restaurantId).list()
    }

}