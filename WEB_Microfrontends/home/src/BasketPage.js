import React, {useRef, useState} from 'react';
import { Toast } from "primereact/toast";
import {useBasket} from "./context/BasketContext";
import BasketItems from "basket/BasketItems";

export default function BasketPage() {
    const { dispatch: basketDispatch, state: basketState } = useBasket();
    const toast = useRef(null);
    const [loading, setLoading] = useState(false);

    const load = () => {
        setLoading(true);

        setTimeout(() => {
            setLoading(false);
        }, 2000);
    };

    const removeFromBasket = (product, index) => {
        basketDispatch({
            type: 'REMOVE_FROM_BASKET',
            payload: index
        });
        toast.current.show({severity:'success', summary: 'Success', detail:`${product.name} removed from basket.`, life: 2500});
    }

    const clearBasket = () => {
        basketDispatch({
            type: 'CLEAR_BASKET'
        });
    }

    return (
        <>
            <Toast ref={toast}/>
            <div className="card">
                <BasketItems items={basketState.items} onRemoveFromBasket={removeFromBasket} onOrderSubmitted={clearBasket}/>
            </div>
        </>
    );
}
