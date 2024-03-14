package com.blazc.test

import com.blazc.OrderGrpc
import com.blazc.OrderServiceGrpc
import io.grpc.ManagedChannelBuilder
import io.quarkus.test.junit.QuarkusTest
import org.junit.jupiter.api.*
import org.junit.jupiter.api.Assertions.assertEquals

@QuarkusTest
class OrderServiceTest {

    companion object {

        lateinit var order: OrderGrpc.Order

        @JvmStatic
        @BeforeAll
        fun setup() {
            order = OrderGrpc.Order.newBuilder()
                .setId("65f33d0f5cc012d6bef8c985")
                .setOrderDate("2024-03-14T17:05:24+00:00")
                .setStatus(OrderGrpc.OrderStatus.PENDING)
                .setAddress("123 Main St")
                .setCustomerName("John Doe")
                .setSellerId("65f33d4419b38f77e27c18e4")
                .setDeliveryPersonId("65f33d4c7bb59d879c2c6f39")
                .setPaymentType(OrderGrpc.PaymentType.CASH)
                .addItems(OrderGrpc.OrderItem.newBuilder()
                    .setId("65f33d0f5cc012d6bef8c985")
                    .setPrice(3000)
                    .setQuantity(2)
                    .setProduct(OrderGrpc.Product.newBuilder()
                        .setId("65f33d7080f352f8900547df")
                        .setName("Pizza Margherita")
                        .setPrice(1500)
                        .build()
                    )
                    .build())
                .setTotalPrice(3000)
                .build()
        }
    }

    @Test
    fun testCreateOrder() {
        val channel = ManagedChannelBuilder.forAddress("localhost", 9001)
            .usePlaintext()
            .build()

        val stub = OrderServiceGrpc.newBlockingStub(channel)

        stub.createOrder(order).let {
            assertEquals("", it.error)
            assertEquals("created", it.message)
        }

        channel.shutdown()
    }
}