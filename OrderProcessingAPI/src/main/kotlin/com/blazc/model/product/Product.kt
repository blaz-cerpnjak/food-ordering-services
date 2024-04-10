package com.blazc.model.product

import com.blazc.OrderGrpc
import org.bson.types.ObjectId


data class Product(
    var id: ObjectId? = null,
    var name: String,
    var price: Int, // 100 = 1
    var image: String
) {

    constructor() : this(
        id = ObjectId(),
        name = "",
        price = 0,
        image = ""
    )

    companion object {
        fun fromGrpc(productGrpc: OrderGrpc.Product): Product {
            return Product(
                id = ObjectId(productGrpc.id),
                name = productGrpc.name,
                price = productGrpc.price,
                image = productGrpc.image
            )
        }

        fun toGrpc(product: Product): OrderGrpc.Product {
            return OrderGrpc.Product.newBuilder()
                .setId(product.id.toString())
                .setName(product.name)
                .setPrice(product.price)
                .setImage(product.image)
                .build()
        }
    }
}
