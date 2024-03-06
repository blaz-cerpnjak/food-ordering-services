package com.blazc.usermanagementapi.rest

import com.blazc.usermanagementapi.dao.UserRepository
import com.blazc.usermanagementapi.vao.User
import org.bson.types.ObjectId
import org.springframework.http.HttpStatus
import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.*
import java.util.logging.Logger


@RestController
@CrossOrigin
class UserController(private val dao: UserRepository) {

    companion object {
        private val log = Logger.getLogger(UserController::class.java.toString())
    }

    @GetMapping("/users")
    fun getAllUsers(): Iterable<com.blazc.usermanagementapi.dto.User> {
        return User.toDtoList(dao.findAll())
    }

    @GetMapping("/users/{id}")
    fun getUserById(@PathVariable id: String): ResponseEntity<com.blazc.usermanagementapi.dto.User> {
        val user = dao.findById(ObjectId(id))
        if (user.isEmpty) {
            log.info { "/users/$id ; User not found!" }
            return ResponseEntity(HttpStatus.NOT_FOUND)
        }
        return ResponseEntity.ok(user.get().toDto())
    }

    @PostMapping("/users")
    fun insertUser(@RequestBody user: com.blazc.usermanagementapi.dto.post.PostUser): ResponseEntity<com.blazc.usermanagementapi.dto.User> {
        val newUser = dao.insert(User(user))
        return ResponseEntity.ok(newUser.toDto())
    }

    @PutMapping("/users/{id}")
    fun updateUser(@PathVariable id: String, @RequestBody user: com.blazc.usermanagementapi.dto.User): ResponseEntity<com.blazc.usermanagementapi.dto.User> {
        val existingUser = dao.findById(ObjectId(id))
        if (existingUser.isEmpty) {
            log.info {"/users/$id ; User not found!" }
            return ResponseEntity(HttpStatus.NOT_FOUND)
        }

        val vao = existingUser.get()
        vao.updateFrom(user)

        dao.save(vao)
        return ResponseEntity.ok(vao.toDto())
    }

    @DeleteMapping("/users/{id}")
    fun deleteUser(@PathVariable id: String): ResponseEntity<String> {
        val existingUser = dao.findById(ObjectId(id))
        if (existingUser.isEmpty) {
            log.info {"/users/$id ; User not found!" }
            return ResponseEntity(HttpStatus.NOT_FOUND)
        }

        val vao = existingUser.get()
        dao.delete(vao)
        return ResponseEntity.ok("deleted")
    }
}