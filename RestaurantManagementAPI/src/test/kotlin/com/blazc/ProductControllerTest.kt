package com.blazc

import com.blazc.model.Product
import io.quarkus.test.junit.QuarkusTest
import io.restassured.RestAssured.given
import jakarta.json.Json
import org.bson.types.ObjectId
import org.hamcrest.CoreMatchers.`is`
import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.Test


@QuarkusTest
class ProductControllerTest {

    private val restaurantId: ObjectId = ObjectId.get()
    private lateinit var createdProduct: Product

    @BeforeEach
    fun createProduct() {
        val productJson = Json.createObjectBuilder()
            .add("id", ObjectId.get().toString())
            .add("restaurantId", restaurantId.toString())
            .add("name", "Hamburger")
            .add("price", 599)
            .build()
        
        val response = given()
            .contentType("application/json")
            .body(productJson.toString())
            .`when`().post("/product")
            .then()
            .statusCode(200)
            .contentType("application/json")
            .body("name", `is`("Hamburger"))
            .extract()
            .response()
    
        createdProduct = response.body.`as`(Product::class.java)
    }
    
    @Test
    fun testUpdateProduct() {
        val productJson = Json.createObjectBuilder()
            .add("id", createdProduct.id.toString())
            .add("restaurantId", createdProduct.restaurantId.toString())
            .add("name", "Cheeseburger")
            .add("price", 699)
            .build()
        
        given()
            .contentType("application/json")
            .body(productJson.toString())
            .`when`().put("/product")
            .then()
            .statusCode(200)
            .contentType("application/json")
            .body("name", `is`("Cheeseburger"))
            .body("price", `is`(699))
    }
    
    @Test
    fun testGetProductById() {
        given()
            .contentType("application/json")
            .`when`().get("/product/${createdProduct.id}")
            .then()
            .statusCode(200)
            .contentType("application/json")
            .body("id", `is`(createdProduct.id.toString()))
    }
    
    @Test
    fun testDeleteProduct() {
        given()
            .`when`().delete("/product/${createdProduct.id}")
            .then()
            .statusCode(200)
    }
}