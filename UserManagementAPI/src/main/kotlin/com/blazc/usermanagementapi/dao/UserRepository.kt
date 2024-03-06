package com.blazc.usermanagementapi.dao

import com.blazc.usermanagementapi.vao.User
import org.bson.types.ObjectId
import org.springframework.data.mongodb.repository.MongoRepository

interface UserRepository : MongoRepository<User, ObjectId> {

}