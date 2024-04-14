import React from "react";
import ReactDOM from "react-dom";

import 'primeflex/primeflex.css';
import 'primereact/resources/primereact.css';
import 'primereact/resources/themes/lara-light-blue/theme.css';
import 'primeicons/primeicons.css';
import "./index.css";

const App = () => (
    <div className="container">
        <div className="mt-10 text-3xl mx-auto max-w-6xl">
            <div>Name: basket</div>
            <div>Framework: react</div>
            <div>Language: JavaScript</div>
            <div>CSS: Tailwind</div>
        </div>
    </div>
);
ReactDOM.render(<App/>, document.getElementById("app"));
