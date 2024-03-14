package com.blazc.model.order

import com.blazc.OrderGrpc

enum class OrderStatus {
    PENDING,
    PREPARING,
    READY_FOR_PICKUP,
    DELIVERING,
    DELIVERED,
    CANCELLED;

    companion object {
        fun fromGrpc(status: OrderGrpc.OrderStatus): OrderStatus {
            return when (status) {
                OrderGrpc.OrderStatus.PENDING -> PENDING
                OrderGrpc.OrderStatus.PREPARING -> PREPARING
                OrderGrpc.OrderStatus.READY_FOR_PICKUP -> READY_FOR_PICKUP
                OrderGrpc.OrderStatus.DELIVERING -> DELIVERING
                OrderGrpc.OrderStatus.DELIVERED -> DELIVERED
                OrderGrpc.OrderStatus.CANCELLED -> CANCELLED
                OrderGrpc.OrderStatus.UNRECOGNIZED -> PENDING
            }
        }

        fun toGrpc(status: OrderStatus): OrderGrpc.OrderStatus {
            return when (status) {
                PENDING -> OrderGrpc.OrderStatus.PENDING
                PREPARING -> OrderGrpc.OrderStatus.PREPARING
                READY_FOR_PICKUP -> OrderGrpc.OrderStatus.READY_FOR_PICKUP
                DELIVERING -> OrderGrpc.OrderStatus.DELIVERING
                DELIVERED -> OrderGrpc.OrderStatus.DELIVERED
                CANCELLED -> OrderGrpc.OrderStatus.CANCELLED
            }
        }
    }
}