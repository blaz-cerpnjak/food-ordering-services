package com.blazc

import com.blazc.model.Restaurant
import io.quarkus.test.junit.QuarkusTest
import io.restassured.RestAssured
import jakarta.json.Json
import org.bson.types.ObjectId
import org.hamcrest.CoreMatchers
import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.Test

@QuarkusTest
class RestaurantControllerTest {

    private lateinit var createdRestaurant: Restaurant

    @BeforeEach
    fun createRestaurant() {
        val json = Json.createObjectBuilder()
            .add("id", ObjectId.get().toString())
            .add("name", "Burger King")
            .add("address", "1234 Main St")
            .add("contact", "555-555-5555")
            .build()

        val response = RestAssured.given()
            .contentType("application/json")
            .body(json.toString())
            .`when`().post("/restaurant")
            .then()
            .statusCode(200)
            .contentType("application/json")
            .body("name", CoreMatchers.`is`("Burger King"))
            .extract()
            .response()

        createdRestaurant = response.body.`as`(Restaurant::class.java)
    }

    @Test
    fun testUpdateRestaurant() {
        val productJson = Json.createObjectBuilder()
            .add("id", createdRestaurant.id.toString())
            .add("name", "McDonalds")
            .add("address", "1234 Main St")
            .add("contact", "555-555-5556")
            .build()

        RestAssured.given()
            .contentType("application/json")
            .body(productJson.toString())
            .`when`().put("/restaurant")
            .then()
            .statusCode(200)
            .contentType("application/json")
            .body("contact", CoreMatchers.`is`("555-555-5556"))
    }

    @Test
    fun testGetRestaurantById() {
        RestAssured.given()
            .contentType("application/json")
            .`when`().get("/restaurant/${createdRestaurant.id}")
            .then()
            .statusCode(200)
            .contentType("application/json")
            .body("id", CoreMatchers.`is`(createdRestaurant.id.toString()))
    }

    @Test
    fun testDeleteRestaurant() {
        RestAssured.given()
            .`when`().delete("/restaurant/${createdRestaurant.id}")
            .then()
            .statusCode(200)
    }

}