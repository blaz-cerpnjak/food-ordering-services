import React from "react";
import { createRoot } from 'react-dom/client';

import "./index.scss";
import 'primeflex/primeflex.css';
import 'primereact/resources/primereact.css';
import 'primereact/resources/themes/lara-light-blue/theme.css';
import 'primeicons/primeicons.css';

import Counter from "./Counter";

const App = () => (
  <div className="mt-10 text-3xl mx-auto max-w-6xl">
    <div>Name: orders</div>
    <div>Framework: react</div>
    <div>Language: JavaScript</div>
    <div>CSS: Tailwind</div>
    <Counter />
  </div>
);

const container = document.getElementById('app');
const root = createRoot(container);
root.render(<App />);
