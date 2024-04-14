import React, {useEffect, useRef, useState} from 'react';
import {Toast} from "primereact/toast";
import {cancelOrder, getOrders} from "../services/OrdersService";
import OrderItems from "./OrderItems";
import {Accordion, AccordionTab} from "primereact/accordion";
import {Button} from "primereact/button";

export default function OrderHistoryItems() {
    const [orders, setOrders] = useState([]);
    const toast = useRef(null);
    const [loading, setLoading] = useState(false);

    useEffect(() => {
        loadOrders();
    }, []);

    const loadOrders = async () => {
        getOrders().then((data) => {
            setOrders(data);
        });
    }

    const cancelCurrentOrder = (order) => {
        cancelOrder(order).then((data) => {
            if (data) {
                toast.current.show({ severity: 'success', summary: 'Success', detail: 'Order cancelled.', life: 3000 });
                loadOrders();
            } else {
                toast.current.show({ severity: 'error', summary: 'Error', detail: 'Error cancelling order.', life: 3000 });
            }
        }).catch((error) => {
            console.error(error);
            toast.current.show({ severity: 'error', summary: 'Error', detail: 'Error cancelling order.', life: 3000 });
        });
    }

    const formatProductPrice = (priceInCents) => {
        const priceInDollars = priceInCents / 100;
        return `$${priceInDollars.toFixed(2)}`;
    };

    const formatDate = (timestamp) => {
        const date = new Date(timestamp * 1000);
        return date.toLocaleDateString();
    };

    return (
        <>
            <Toast ref={toast}/>
            <div className="card">
                {orders.length > 0 ?
                    <Accordion activeIndex={0}>
                        {orders.map((order, index) => (
                            <AccordionTab key={index} header={`Order - ${formatProductPrice(order.totalPrice)} (${formatDate(order.timestamp)})`}>
                                <OrderItems order={order}/>
                                {order.status === 'PENDING' && <Button icon="pi pi-shopping-cart" label="Cancel Order" onClick={() => cancelCurrentOrder(order)}></Button>}
                            </AccordionTab>
                        ))}
                    </Accordion>
                    :
                    <div className="text-center">
                        <h2>No orders found.</h2>
                    </div>
                }
            </div>
        </>
    );
}
