export const createOrder = (order) =>
    fetch(`/api/v1/orders`, {
        method: 'POSt',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
        },
        body: JSON.stringify(order)
    }).then((res) => {
        if (!res.ok) {
            throw new Error('Error creating order');
        }
        return res.json();
    }).then((data) => {
        console.log(data);
        return data;
    }).catch((error) => {
        console.error(error);
        return null;
    });
