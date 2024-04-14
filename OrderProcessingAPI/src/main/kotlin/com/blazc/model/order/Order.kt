package com.blazc.model.order

import com.blazc.OrderGrpc
import com.blazc.model.payment.PaymentType
import com.blazc.model.product.Product
import org.bson.types.ObjectId
import java.time.Instant

data class Order (
    var id: ObjectId? = null,
    var sellerId: ObjectId? = null,
    var deliveryPersonId: ObjectId? = null,
    var address: String,
    var customerName: String,
    var items: List<Product>,
    var status: OrderStatus,
    var timestamp: Long,
    var paymentType: PaymentType,
    var totalPrice: Int
) {

    constructor() : this(
        id = ObjectId(),
        sellerId = ObjectId(),
        deliveryPersonId = ObjectId(),
        address = "",
        customerName = "",
        items = emptyList(),
        status = OrderStatus.PENDING,
        timestamp = Instant.now().epochSecond,
        paymentType = PaymentType.CASH,
        totalPrice = 0
    )

    companion object {
        fun fromGrpc(order: OrderGrpc.Order): Order {
            return Order(
                id = ObjectId(order.id),
                sellerId = ObjectId(order.sellerId),
                deliveryPersonId = ObjectId(order.deliveryPersonId),
                address = order.address,
                customerName = order.customerName,
                items = order.itemsList.map {
                    Product.fromGrpc(it)
                },
                status = OrderStatus.fromGrpc(order.status),
                timestamp = order.timestamp,
                paymentType = PaymentType.fromGrpc(order.paymentType),
                totalPrice = order.totalPrice
            )
        }

        fun toGrpc(order: Order): OrderGrpc.Order {
            return OrderGrpc.Order.newBuilder()
                .setId(order.id.toString())
                .setSellerId(order.sellerId.toString())
                .setDeliveryPersonId(order.deliveryPersonId.toString())
                .setAddress(order.address)
                .setCustomerName(order.customerName)
                .addAllItems(order.items.map { Product.toGrpc(it) })
                .setStatus(OrderStatus.toGrpc(order.status))
                .setTimestamp(order.timestamp)
                .setPaymentType(PaymentType.toGrpc(order.paymentType))
                .setTotalPrice(order.totalPrice)
                .build()
        }
    }
}
