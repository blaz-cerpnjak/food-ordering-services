export const getProducts = () =>
    fetch(`/api/v1/products`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json',
        },
    }).then((res) => res.json());
