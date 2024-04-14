import React, {useRef, useState} from 'react';
import { Button } from 'primereact/button';
import { Rating } from 'primereact/rating';
import { Tag } from 'primereact/tag';
import { classNames } from 'primereact/utils';
import { Toast } from "primereact/toast";
import { DataView } from 'primereact/dataview';
import { createOrder } from "../services/OrderService";
import ObjectId from 'bson-objectid';

export default function BasketItems({ items, onRemoveFromBasket, onOrderSubmitted }) {
    const toast = useRef(null);
    const [loading, setLoading] = useState(false);

    const submitOrder = async (items) => {
        setLoading(true);

        const currentTimestamp = Math.floor(Date.now() / 1000);

        console.log(currentTimestamp)

        const order = {
            id: ObjectId().toString(),
            address: '1234 Elm Street',
            customerName: 'John Doe',
            items: items,
            timestamp: currentTimestamp,
            paymentType: 'CREDIT_CARD',
            sellerId: '000000000000000000000000',
            deliveryGuyId: '000000000000000000000000',
            status: 'PENDING',
            totalPrice: items.reduce((acc, item) => acc + item.price, 0)
        }

        createOrder(order).then((data) => {
            if (data) {
                toast.current.show({severity:'success', summary: 'Success', detail:'Order submitted.', life: 3000});
                onOrderSubmitted();
            } else {
                toast.current.show({severity:'error', summary: 'Error', detail:'Error submitting order.', life: 3000});
            }
        }).catch((error) => {
            console.error(error);
            setLoading(false)
            toast.current.show({severity:'error', summary: 'Error', detail:'Error submitting order.', life: 3000});
        });
    }

    const formatProductPrice = (priceInCents) => {
        const priceInDollars = priceInCents / 100;
        return `$${priceInDollars.toFixed(2)}`;
    };

    const itemTemplate = (product, index) => {
        return (
            <div className="col-12" key={product.id}>
                <div className={classNames('flex flex-column xl:flex-row xl:align-items-start p-4 gap-4', { 'border-top-1 surface-border': index !== 0 })}>
                    <img className="w-9 sm:w-16rem xl:w-10rem shadow-2 block xl:block mx-auto border-round" src={`${product.image}`} alt={product.name} />
                    <div className="flex flex-column sm:flex-row justify-content-between align-items-center xl:align-items-start flex-1 gap-4">
                        <div className="flex flex-column align-items-center sm:align-items-start gap-3">
                            <div className="text-2xl font-bold text-900">{product.name}</div>
                            <Rating value={product.rating} readOnly cancel={false}></Rating>
                            <div className="flex align-items-center gap-3">
                                <span className="flex align-items-center gap-2">
                                    <i className="pi pi-tag"></i>
                                    <span className="font-semibold">{product.category}</span>
                                </span>
                                <Tag value="IN STOCK" severity="success"></Tag>
                            </div>
                        </div>
                        <div className="flex sm:flex-column align-items-center sm:align-items-end gap-3 sm:gap-2">
                            <span className="text-2xl font-semibold">{formatProductPrice(product.price)}</span>
                            <Button icon="pi pi-trash" className="p-button-rounded" severity="danger" onClick={() => onRemoveFromBasket(product, index)}></Button>
                        </div>
                    </div>
                </div>
            </div>
        );
    };

    const listTemplate = (items) => {
        if (!items || items.length === 0) return null;

        let list = items.map((product, index) => {
            return itemTemplate(product, index);
        });

        return <div className="grid grid-nogutter">{list}</div>;
    };

    return (
        <>
            <Toast ref={toast}/>
            <div className="card">
                {items.length > 0 ?
                    <>
                        <div className="card">
                            <DataView value={items} listTemplate={listTemplate}/>
                        </div>
                        <div className="card flex flex-wrap justify-content-center gap-3">
                            <Button label="Submit Order" icon="pi pi-check" loading={loading} onClick={() => submitOrder(items)}/>
                        </div>
                    </>
                    :
                    <div className="text-center">
                        <h2>Your basket is empty</h2>
                    </div>
                }
            </div>
        </>
    );
}
