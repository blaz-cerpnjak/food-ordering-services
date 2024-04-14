import React from "react";
import {PrimeReactProvider} from "primereact/api";
import {BrowserRouter, Route, Routes} from "react-router-dom";

import "./index.scss";
import 'primeflex/primeflex.css';
import 'primereact/resources/primereact.css';
import 'primereact/resources/themes/lara-light-blue/theme.css';
import 'primeicons/primeicons.css';

import OrderHistoryItems from "orders/OrderHistoryItems";
import Navbar from "./Navbar";
import ProductsPage from "./ProductsPage";
import {createRoot} from "react-dom/client";
import {BasketProvider, useBasket} from "./context/BasketContext";
import BasketPage2 from "./BasketPage";

const App = () => (
    <PrimeReactProvider>
        <BasketProvider>
            <BrowserRouter>
                <div className="App">
                    <Navbar/>
                    <div className="container">
                        <Routes>
                            <Route exact path="/" element={<ProductsPage />} />
                            <Route path="/orders" element={<OrderHistoryItems />} />
                            <Route path="/basket" element={<BasketPage2 />} />
                        </Routes>
                    </div>
                </div>
            </BrowserRouter>
        </BasketProvider>
    </PrimeReactProvider>
);

const container = document.getElementById('app');
const root = createRoot(container);
root.render(<App />);

