package com.blazc.controller

import com.blazc.repository.ProductRepository
import com.blazc.model.Product
import io.smallrye.mutiny.Uni
import jakarta.inject.Inject
import jakarta.ws.rs.*
import jakarta.ws.rs.core.MediaType
import org.bson.types.ObjectId

@Path("/product")
class ProductController {

    @Inject
    lateinit var productRepository: ProductRepository

    @POST
    @Consumes(MediaType.APPLICATION_JSON)
    @Produces(MediaType.APPLICATION_JSON)
    fun createProduct(product: Product): Uni<Product> {
        return productRepository.create(product)
    }

    @GET
    @Path("/{id}")
    @Produces(MediaType.APPLICATION_JSON)
    fun getProductById(@PathParam("id") id: String): Uni<Product> {
        lateinit var objectId: ObjectId

        try {
            objectId = ObjectId(id)
        } catch (e: Exception) {
            return Uni.createFrom().failure(e)
        }

        return productRepository.findById(objectId)
    }

    @GET
    @Produces(MediaType.APPLICATION_JSON)
    fun getAllProducts(): Uni<List<Product>> {
        return productRepository.getAll()
    }

    @GET
    @Path("/restaurant/{id}")
    @Produces(MediaType.APPLICATION_JSON)
    fun getAllProductsByRestaurantId(@PathParam("id") restaurantId: String): Uni<List<Product>> {
        return productRepository.getAllByRestaurantId(restaurantId)
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
        lateinit var objectId: ObjectId

        try {
            objectId = ObjectId(id)
        } catch (e: Exception) {
            return Uni.createFrom().failure(e)
        }

        return productRepository.deleteById(objectId)
    }

}