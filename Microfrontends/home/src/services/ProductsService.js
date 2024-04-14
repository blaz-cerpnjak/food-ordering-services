export const getProducts = () =>
    fetch(`/api/v1/products/restaurant/66163e8b15226a4958756a26`).then((res) => res.json());
