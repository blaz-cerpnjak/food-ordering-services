package com.blazc.repository

import com.blazc.model.Product
import io.quarkus.mongodb.panache.reactive.ReactivePanacheMongoRepository
import io.smallrye.mutiny.Uni
import jakarta.enterprise.context.ApplicationScoped

@ApplicationScoped
class ProductRepository: ReactivePanacheMongoRepository<Product> {

    fun create(product: Product): Uni<Product> {
        return persist(product)
    }

    fun getAllByRestaurantId(restaurantId: String): Uni<List<Product>> {
        return find("restaurantId", restaurantId).list()
    }

    fun getAll(): Uni<List<Product>> {
        return listAll()
    }

}