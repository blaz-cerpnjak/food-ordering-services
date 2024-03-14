package com.blazc.model.order

import com.blazc.OrderGrpc
import com.blazc.model.payment.PaymentType
import com.blazc.utils.DateFormatter
import org.bson.types.ObjectId
import java.time.LocalDateTime

data class Order (
    var id: ObjectId? = null,
    var sellerId: ObjectId? = null,
    var deliveryPersonId: ObjectId? = null,
    var address: String,
    var customerName: String,
    var items: List<OrderItem>,
    var status: OrderStatus,
    var orderDate: LocalDateTime,
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
        orderDate = LocalDateTime.now(),
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
                    OrderItem.fromGrpc(it)
                },
                status = OrderStatus.fromGrpc(order.status),
                orderDate = DateFormatter.toLocalDateTime(order.orderDate),
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
                .addAllItems(order.items.map {
                    OrderItem.toGrpc(it)
                })
                .setStatus(OrderStatus.toGrpc(order.status))
                .setOrderDate(DateFormatter.toString(order.orderDate))
                .setPaymentType(PaymentType.toGrpc(order.paymentType))
                .setTotalPrice(order.totalPrice)
                .build()
        }
    }
}
