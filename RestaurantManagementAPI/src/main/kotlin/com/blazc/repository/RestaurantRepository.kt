package com.blazc.repository

import com.blazc.model.Restaurant
import io.quarkus.mongodb.panache.reactive.ReactivePanacheMongoRepository
import io.smallrye.mutiny.Uni
import jakarta.enterprise.context.ApplicationScoped

@ApplicationScoped
class RestaurantRepository: ReactivePanacheMongoRepository<Restaurant> {

    fun add(restaurant: Restaurant): Uni<Restaurant> {
        return persist(restaurant)
    }

    fun findByName(name: String): Uni<Restaurant?> {
        return find("name", name).firstResult()
    }

    fun getAll(): Uni<List<Restaurant>> {
        return listAll()
    }

}