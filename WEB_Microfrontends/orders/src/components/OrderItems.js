import React, {useRef} from 'react';
import { Rating } from 'primereact/rating';
import { Tag } from 'primereact/tag';
import { classNames } from 'primereact/utils';
import { Toast } from "primereact/toast";
import { DataView } from 'primereact/dataview';

export default function OrderItems({ order }) {
    const toast = useRef(null);

    const formatProductPrice = (priceInCents) => {
        const priceInDollars = priceInCents / 100;
        return `$${priceInDollars.toFixed(2)}`;
    };

    const getSeverity = (status) => {
        switch (status) {
            case 'PENDING':
                return 'warning';
            case 'DELIVERED':
                return 'success';
            case 'CANCELLED':
                return 'danger';
            default:
                return 'info';
        }
    }

    const itemTemplate = (product, index) => {
        return (
            <div className="col-12" key={product.id}>
                <div className={classNames('flex flex-column xl:flex-row xl:align-items-start p-4 gap-4', { 'border-top-1 surface-border': index !== 0 })}>
                    <img className="w-9 sm:w-16rem xl:w-10rem shadow-2 block xl:block mx-auto border-round" src={`${product.image}`} alt={product.name} />
                    <div
                        className="flex flex-column sm:flex-row justify-content-between align-items-center xl:align-items-start flex-1 gap-4">
                        <div className="flex flex-column align-items-center sm:align-items-start gap-3">
                            <div className="text-2xl font-bold text-900">{product.name}</div>
                            <Rating value={product.rating} readOnly cancel={false}></Rating>
                            <div className="flex align-items-center gap-3">
                                <span className="flex align-items-center gap-2">
                                    <i className="pi pi-tag"></i>
                                    <span className="font-semibold">{product.category}</span>
                                </span>
                                <Tag value={order.status} severity={getSeverity(order.status)}></Tag>
                            </div>
                        </div>
                        <div className="flex sm:flex-column align-items-center sm:align-items-end gap-3 sm:gap-2">
                            <span className="text-2xl font-semibold">{formatProductPrice(product.price)}</span>
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
            <div>
                <Toast ref={toast} />
                {order.items.length > 0 ?
                    <>
                    <div className="card">
                            <DataView value={order.items} listTemplate={listTemplate}/>
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
