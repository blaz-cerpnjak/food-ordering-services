package com.blazc.usermanagementapi

import com.blazc.usermanagementapi.dto.User
import com.blazc.usermanagementapi.dto.post.PostUser
import com.blazc.usermanagementapi.rest.UserController
import com.mongodb.assertions.Assertions.assertTrue
import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.Test
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.test.context.ActiveProfiles

@SpringBootTest
@ActiveProfiles("test")
class UserManagementApiApplicationTests {

    @Autowired
    lateinit var userController: UserController

    lateinit var createdUser: User

    @BeforeEach
    fun setup() {
        val mockUser = PostUser(
            "John",
            "Doe",
            "john@example.com",
            "1234",
            "1234567890",
            "buyer"
        )

        val response = userController.insertUser(mockUser)
        assertTrue(response.body != null)

        createdUser = response.body!!
    }

    @Test
    fun testGetUserById() {
        val user = userController.getUserById(createdUser.id)
        assertTrue(user.body != null)
        assertEquals(createdUser, user.body)
    }

    @Test
    fun updateUser() {
        val updatedUser = createdUser.copy(name = "Jane")
        val response = userController.updateUser(createdUser.id, updatedUser)
        assertTrue(response.body != null)
        assertEquals(updatedUser, response.body)
    }

    @Test
    fun testDeleteUser() {
        val response = userController.deleteUser(createdUser.id)
        assertEquals("deleted", response.body)

        val user = userController.getUserById(createdUser.id)
        assertTrue(user.body == null)
    }
}
