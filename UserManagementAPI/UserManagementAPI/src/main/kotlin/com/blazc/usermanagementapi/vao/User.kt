package com.blazc.usermanagementapi.vao

import com.blazc.usermanagementapi.dto.User
import org.bson.types.ObjectId
import org.springframework.data.annotation.Id
import org.springframework.data.mongodb.core.mapping.Document

@Document(collection = "products")
class User {
    @Id
    var id: ObjectId
    var name: String
    var lastname: String
    var email: String
    var password: String
    var phone: String
    var role: String // "buyer", "seller", "delivery"

    companion object {
        fun toDtoList(users: Iterable<com.blazc.usermanagementapi.vao.User>): List<User> {
            val ret = ArrayList<User>()
            for (user in users) {
                ret.add(user.toDto())
            }
            return ret
        }
    }

    constructor(dto: com.blazc.usermanagementapi.dto.User) {
        this.id = dto.id
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

    fun updateFrom(dto: com.blazc.usermanagementapi.dto.User) {
        this.id = dto.id
        this.name = dto.name
        this.lastname = dto.lastname
        this.email = dto.email
        this.password = dto.password
        this.phone = dto.phone
        this.role = dto.role
    }

    fun toDto(): com.blazc.usermanagementapi.dto.User {
        return com.blazc.usermanagementapi.dto.User(
                this.id,
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