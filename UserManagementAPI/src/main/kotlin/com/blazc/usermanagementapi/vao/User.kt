package com.blazc.usermanagementapi.vao

import com.blazc.usermanagementapi.dto.User
import org.bson.types.ObjectId
import org.springframework.data.annotation.Id
import org.springframework.data.mongodb.core.mapping.Document
import java.util.UUID

@Document(collection = "products")
class User {
    @Id
    var id: ObjectId = ObjectId.get()
    var name: String = ""
    var lastname: String = ""
    var email: String = ""
    var password: String = ""
    var phone: String = ""
    var role: String = "" // "buyer", "seller", "delivery"

    companion object {
        fun toDtoList(users: Iterable<com.blazc.usermanagementapi.vao.User>): List<User> {
            val ret = ArrayList<User>()
            for (user in users) {
                ret.add(user.toDto())
            }
            return ret
        }
    }

    constructor() {}

    constructor(dto: User) {
        this.id = ObjectId(dto.id)
        this.name = dto.name
        this.lastname = dto.lastname
        this.email = dto.email
        this.password = dto.password
        this.phone = dto.phone
        this.role = dto.role
    }

    constructor(user: com.blazc.usermanagementapi.dto.post.PostUser) {
        this.id = ObjectId.get()
        this.name = user.name
        this.lastname = user.lastname
        this.email = user.email
        this.password = user.password
        this.phone = user.phone
        this.role = user.role
    }

    fun updateFrom(dto: User) {
        this.id = ObjectId(dto.id)
        this.name = dto.name
        this.lastname = dto.lastname
        this.email = dto.email
        this.password = dto.password
        this.phone = dto.phone
        this.role = dto.role
    }

    fun toDto(): User {
        return User(
                this.id.toString(),
                this.name,
                this.lastname,
                this.email,
                this.password,
                this.phone,
                this.role
        )
    }

    override fun toString(): String {
        return "User (id=$id, name='$name', lastname='$lastname', email='$email', password='$password', phone='$phone', role='$role')"
    }
}