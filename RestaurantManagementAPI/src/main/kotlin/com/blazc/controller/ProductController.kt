package com.blazc.controller

import com.blazc.model.Product
import com.blazc.repository.ProductRepository
import io.smallrye.mutiny.Uni
import jakarta.inject.Inject
import jakarta.ws.rs.*
import jakarta.ws.rs.core.MediaType
import org.bson.types.ObjectId
import org.jboss.logging.Logger

@Path("/product")
class ProductController {

    companion object {
        val LOG: Logger = Logger.getLogger(ProductController::class.java)
    }

    @Inject
    lateinit var productRepository: ProductRepository

    @POST
    @Consumes(MediaType.APPLICATION_JSON)
    @Produces(MediaType.APPLICATION_JSON)
    fun createProduct(product: Product): Uni<Product> {
        return productRepository.create(product)
    }

    @GET
    @Produces(MediaType.APPLICATION_JSON)
    fun getAllProducts(): Uni<List<Product>> {
        return productRepository.listAll()
    }

    @GET
    @Path("/{id}")
    @Produces(MediaType.APPLICATION_JSON)
    fun getProductById(@PathParam("id") id: String): Uni<Product> {
        val objectId: ObjectId

        try {
            objectId = ObjectId(id)
        } catch (e: Exception) {
            LOG.error("Error parsing id", e)
            return Uni.createFrom().failure(e)
        }

        return productRepository.findById(objectId)
    }

    @GET
    @Path("/restaurant/{id}")
    @Produces(MediaType.APPLICATION_JSON)
    fun getAllProductsByRestaurantId(@PathParam("id") restaurantId: String): Uni<List<Product>> {
        val objectId: ObjectId

        try {
            objectId = ObjectId(restaurantId)
        } catch (e: Exception) {
            LOG.error("Error parsing id", e)
            return Uni.createFrom().failure(e)
        }

        return productRepository.getAllByRestaurantId(objectId)
    }


    @PUT
    @Consumes(MediaType.APPLICATION_JSON)
    @Produces(MediaType.APPLICATION_JSON)
    fun updateProduct(product: Product): Uni<Product> {
        return productRepository.update(product)
    }

    @DELETE
    @Path("/{id}")
    @Produces(MediaType.APPLICATION_JSON)
    fun deleteProduct(@PathParam("id") id: String): Uni<Boolean> {
        val objectId: ObjectId

        try {
            objectId = ObjectId(id)
        } catch (e: Exception) {
            LOG.error("Error parsing id", e)
            return Uni.createFrom().failure(e)
        }

        return productRepository.deleteById(objectId)
    }

}