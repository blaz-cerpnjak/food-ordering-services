package com.blazc.service

import com.blazc.OrderGrpc
import com.blazc.OrderServiceGrpc
import com.blazc.model.order.Order
import com.blazc.repository.OrderRepository
import io.grpc.Status
import io.grpc.stub.StreamObserver
import io.quarkus.grpc.GrpcService
import jakarta.inject.Inject
import jakarta.inject.Singleton
import org.bson.types.ObjectId
import org.eclipse.microprofile.reactive.messaging.Channel
import org.eclipse.microprofile.reactive.messaging.Emitter

@GrpcService
@Singleton
class OrderService : OrderServiceGrpc.OrderServiceImplBase() {

    @Inject
    lateinit var orderRepository: OrderRepository

    //@Channel("orders")
    //private lateinit var emmiter: Emitter<String>

    override fun health(request: OrderGrpc.Empty, responseObserver: StreamObserver<OrderGrpc.Confirmation>) {
        val confirmation = OrderGrpc.Confirmation.newBuilder()
            .setError("")
            .setMessage("success")
            .build()

        responseObserver.onNext(confirmation)
        responseObserver.onCompleted()
    }

    override fun createOrder(request: OrderGrpc.Order, responseObserver: StreamObserver<OrderGrpc.Confirmation>) {
        val order: Order?

        try {
            order = Order.fromGrpc(request)
        } catch (e: Exception) {
            e.printStackTrace()
            responseObserver.onError(Status.INVALID_ARGUMENT.withDescription("invalid order format").asRuntimeException())
            responseObserver.onCompleted()
            return
        }

        orderRepository.persistOrUpdate(order)

        val response = OrderGrpc.Confirmation.newBuilder()
            .setError("")
            .setMessage("created")
            .build()

        //emmiter.send("${order.id}:${order.status}")

        responseObserver.onNext(response)
        responseObserver.onCompleted()
    }

    override fun getOrder(request: OrderGrpc.GetOrderRequest, responseObserver: StreamObserver<OrderGrpc.Order>) {
        val orderId: ObjectId?

        try {
            orderId = ObjectId(request.id)
        } catch (e: Exception) {
            responseObserver.onError(Status.INVALID_ARGUMENT.withDescription("invalid objectId format").asRuntimeException())
            responseObserver.onCompleted()
            return
        }

        println(orderId)

        val order = orderRepository.findById(orderId)
        if (order == null) {
            print("not found")
            responseObserver.onError(Status.NOT_FOUND.asRuntimeException())
            responseObserver.onCompleted()
            return
        }

        println("found")

        val response = Order.toGrpc(order)
        responseObserver.onNext(response)
        responseObserver.onCompleted()
    }

    override fun getOrders(request: OrderGrpc.Empty?, responseObserver: StreamObserver<OrderGrpc.Order>) {
        val orders = orderRepository.listAll()

        orders.forEach {
            responseObserver.onNext(Order.toGrpc(it))
        }

        responseObserver.onCompleted()
    }

    override fun getOrdersBySeller(request: OrderGrpc.GetOrdersRequest, responseObserver: StreamObserver<OrderGrpc.Order>) {
        val sellerId: ObjectId?

        try {
            sellerId = ObjectId(request.id)
        } catch (e: Exception) {
            responseObserver.onError(Status.INVALID_ARGUMENT.withDescription("invalid objectId format").asRuntimeException())
            responseObserver.onCompleted()
            return
        }

        val orders = orderRepository.findAllBySellerId(sellerId)

        orders.forEach {
            responseObserver.onNext(Order.toGrpc(it))
        }

        responseObserver.onCompleted()
    }

    override fun getOrdersByDeliveryPerson(request: OrderGrpc.GetOrdersRequest, responseObserver: StreamObserver<OrderGrpc.Order>) {
        val deliveryPersonId: ObjectId?

        try {
            deliveryPersonId = ObjectId(request.id)
        } catch (e: Exception) {
            responseObserver.onError(Status.INVALID_ARGUMENT.withDescription("invalid objectId format").asRuntimeException())
            responseObserver.onCompleted()
            return
        }

        val orders = orderRepository.findAllByDeliveryPersonId(deliveryPersonId)

        print(orders.size)

        orders.forEach {
            responseObserver.onNext(Order.toGrpc(it))
        }

        responseObserver.onCompleted()
    }

    override fun updateOrder(request: OrderGrpc.Order, responseObserver: StreamObserver<OrderGrpc.Confirmation>) {
        val order: Order?

        try {
            order = Order.fromGrpc(request)
        } catch (e: Exception) {
            responseObserver.onError(Status.INVALID_ARGUMENT.withDescription("invalid order format").asRuntimeException())
            responseObserver.onCompleted()
            return
        }

        orderRepository.update(order)

        val response = OrderGrpc.Confirmation.newBuilder()
            .setError("")
            .setMessage("updated")
            .build()

        //emmiter.send("${order.id}:${order.status}")

        responseObserver.onNext(response)
        responseObserver.onCompleted()
    }

    override fun deleteOrder(request: OrderGrpc.DeleteOrderRequest, responseObserver: StreamObserver<OrderGrpc.Confirmation>) {
        val orderId: ObjectId?

        try {
            orderId = ObjectId(request.id)
        } catch (e: Exception) {
            responseObserver.onError(Status.INVALID_ARGUMENT.withDescription("invalid objectId format").asRuntimeException())
            responseObserver.onCompleted()
            return
        }

        val deleted = orderRepository.deleteById(orderId)

        val response = OrderGrpc.Confirmation.newBuilder()
            .setError(if (deleted) "" else "order was not found")
            .setMessage(if (deleted) "deleted" else "error while deleting")
            .build()

        responseObserver.onNext(response)
        responseObserver.onCompleted()
    }
}