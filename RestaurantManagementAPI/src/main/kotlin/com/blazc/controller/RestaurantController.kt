package com.blazc.controller

import com.blazc.model.Restaurant
import com.blazc.repository.RestaurantRepository
import io.smallrye.mutiny.Uni
import jakarta.inject.Inject
import jakarta.ws.rs.*
import jakarta.ws.rs.core.MediaType
import org.bson.types.ObjectId
import org.jboss.logging.Logger

@Path("/restaurant")
class RestaurantController {

    companion object {
        val LOG: Logger = Logger.getLogger(RestaurantController::class.java)
    }

    @Inject
    lateinit var restaurantRepository: RestaurantRepository

    @POST
    @Consumes(MediaType.APPLICATION_JSON)
    @Produces(MediaType.APPLICATION_JSON)
    fun createRestaurant(restaurant: Restaurant): Uni<Restaurant> {
        return restaurantRepository.add(restaurant)
    }

    @GET
    @Path("/{id}")
    @Produces(MediaType.APPLICATION_JSON)
    fun getRestaurantById(@PathParam("id") id: String): Uni<Restaurant?> {
        lateinit var objectId: ObjectId

        try {
            objectId = ObjectId(id)
        } catch (e: Exception) {
            LOG.error("Error parsing id", e)
            return Uni.createFrom().failure(e)
        }

        return restaurantRepository.findById(objectId)
    }

    @GET
    @Produces(MediaType.APPLICATION_JSON)
    fun getAllRestaurants(): Uni<List<Restaurant>> {
        return restaurantRepository.getAll()
    }

    @PUT
    @Produces(MediaType.APPLICATION_JSON)
    fun updateRestaurant(restaurant: Restaurant): Uni<Restaurant> {
        return restaurantRepository.update(restaurant)
    }

    @DELETE
    @Path("/{id}")
    @Produces(MediaType.APPLICATION_JSON)
    fun deleteRestaurant(@PathParam("id") id: String): Uni<Boolean> {
        lateinit var objectId: ObjectId

        try {
            objectId = ObjectId(id)
        } catch (e: Exception) {
            LOG.error("Error parsing id", e)
            return Uni.createFrom().failure(e)
        }

        return restaurantRepository.deleteById(objectId)
    }

}