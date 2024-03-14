package com.blazc.model.order

import com.blazc.OrderGrpc
import com.blazc.model.product.Product
import org.bson.types.ObjectId

data class OrderItem(
    var id: ObjectId? = null,
    var product: Product,
    var quantity: Int,
    var price: Int
) {

    constructor() : this(
        id = ObjectId(),
        product = Product(),
        quantity = 0,
        price = 0
    )

    companion object {
        fun fromGrpc(orderItem: OrderGrpc.OrderItem): OrderItem {
            return OrderItem(
                id = ObjectId(orderItem.id),
                product = Product.fromGrpc(orderItem.product),
                quantity = orderItem.quantity,
                price = orderItem.price
            )
        }

        fun toGrpc(orderItem: OrderItem): OrderGrpc.OrderItem {
            return OrderGrpc.OrderItem.newBuilder()
                .setId(orderItem.id.toString())
                .setProduct(Product.toGrpc(orderItem.product))
                .setQuantity(orderItem.quantity)
                .setPrice(orderItem.price)
                .build()
        }
    }
}
