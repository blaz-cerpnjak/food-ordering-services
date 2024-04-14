import React from 'react';
import { Menubar } from 'primereact/menubar';
import { Badge } from 'primereact/badge';
import { Avatar } from 'primereact/avatar';
import { Button } from 'primereact/button';
import { Link } from 'react-router-dom';
import { useBasket } from "./context/BasketContext";

export default function Navbar() {
    const { state: basketState } = useBasket();
    const badgeNumber = basketState.items.length;

    const itemRenderer = (item) => (
        <Link to={item.path} className="flex align-items-center p-menuitem-link">
            <span className={item.icon} />
            <span className="mx-2">{item.label}</span>
            {(item.badge && item.badge !== 0) ? <Badge value={item.badge} /> : null}
        </Link>
    );
    const items = [
        {
            label: 'Home',
            icon: 'pi pi-home',
            path: '/',
            template: itemRenderer
        },
        {
            label: 'Orders',
            icon: 'pi pi-envelope',
            path: '/orders',
            badge: 3,
            template: itemRenderer
        },
        {
            label: 'Basket',
            icon: 'pi pi-shopping-cart',
            path: '/basket',
            badge: badgeNumber,
            template: itemRenderer
        }
    ];

    const start = <img alt="logo" src="https://primefaces.org/cdn/primereact/images/logo.png" height="40" className="mr-2"></img>;
    const end = (
        <div className="flex align-items-center gap-2">
            <Link to="/basket">
                <Button icon="pi pi-shopping-cart" rounded text severity="secondary" aria-label="Filter" />
            </Link>
            <Avatar image="https://cdn.pixabay.com/photo/2018/11/08/23/52/man-3803551_1280.jpg" shape="circle" />
        </div>
    );

    return (
        <div className="card">
            <Menubar model={items} start={start} end={end} />
        </div>
    )
}
