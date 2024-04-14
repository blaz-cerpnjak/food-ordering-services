import React, { createContext, useContext, useReducer } from 'react';

// Define the initial state for the basket
const initialState = {
    items: [],
};

// Define the actions to modify the basket
const ADD_TO_BASKET = 'ADD_TO_BASKET';
const REMOVE_FROM_BASKET = 'REMOVE_FROM_BASKET';
const CLEAR_BASKET = 'CLEAR_BASKET';

// Define a reducer function to handle state changes
function basketReducer(state, action) {
    switch (action.type) {
        case ADD_TO_BASKET:
            return {
                ...state,
                items: [...state.items, action.payload],
            };
        case REMOVE_FROM_BASKET:
            return {
                ...state,
                items: state.items.filter((_, index) => index !== action.payload),
            };
        case CLEAR_BASKET:
            return {
                ...state,
                items: [],
            };
        default:
            return state;
    }
}

// Create a context for the basket
const BasketContext = createContext();

// Create a provider component to manage the basket state
export function BasketProvider({ children }) {
    const [state, dispatch] = useReducer(basketReducer, initialState);

    return (
        <BasketContext.Provider value={{ state, dispatch }}>
            {children}
        </BasketContext.Provider>
    );
}

// Custom hook to use the basket context
export function useBasket() {
    return useContext(BasketContext);
}
