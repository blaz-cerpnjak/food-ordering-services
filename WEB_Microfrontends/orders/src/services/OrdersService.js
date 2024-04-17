export const getOrders = (order) =>
    fetch(`/api/v1/orders`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
            'Authorization': `Bearer ${process.env.API_TOKEN}`,
        }
    }).then((res) => {
        if (!res.ok) {
            throw new Error('Error fetching orders');
        }
        return res.json();
    })

export const cancelOrder = (order) =>
    fetch(`/api/v1/orders/${order.id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
            'Authorization': `Bearer ${process.env.API_TOKEN}`,
        },
        body: JSON.stringify({
            ...order,
            status: 'CANCELLED',
        })
    }).then((res) => {
        if (!res.ok) {
            throw new Error('Error cancelling order');
        }
        return res.json();
    })
