package com.blazc.model.payment

import com.blazc.OrderGrpc

enum class PaymentType {
    CREDIT_CARD,
    CASH;

    companion object {
        fun fromGrpc(type: OrderGrpc.PaymentType): PaymentType {
            return when (type) {
                OrderGrpc.PaymentType.CREDIT_CARD -> PaymentType.CREDIT_CARD
                OrderGrpc.PaymentType.CASH -> PaymentType.CASH
                OrderGrpc.PaymentType.UNRECOGNIZED -> PaymentType.CASH
            }
        }

        fun toGrpc(type: PaymentType): OrderGrpc.PaymentType {
            return when (type) {
                PaymentType.CREDIT_CARD -> OrderGrpc.PaymentType.CREDIT_CARD
                PaymentType.CASH -> OrderGrpc.PaymentType.CASH
            }
        }
    }
}