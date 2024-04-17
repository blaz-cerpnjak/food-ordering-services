export const getProducts = () =>
    fetch(`/api/v1/products`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
            'Authorization': `Bearer ${process.env.API_TOKEN}`,
        },
    }).then((res) => res.json())
