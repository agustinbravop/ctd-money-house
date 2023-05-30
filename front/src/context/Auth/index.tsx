/* eslint-disable @typescript-eslint/no-empty-function */
import React, { createContext, SetStateAction, useState } from 'react';
import { useLocalStorage } from '../../hooks';

export const AuthContext = createContext<{
    isAuthenticated: boolean;
    setIsAuthenticated: React.Dispatch<SetStateAction<boolean>>;
    logout: () => void;
}>({
    isAuthenticated: false,
    setIsAuthenticated: () => {
    },
    logout: () => {
    },
});

const AuthProvider = ({ children }: { children: React.ReactNode }) => {
    const [token, setToken] = useLocalStorage('token');
    const [isAuthenticated, setIsAuthenticated] = useState(!!token);

    const logout = () => {
        setIsAuthenticated(false);
        setToken(null);
    };

    return (
        <AuthContext.Provider
            value={{ isAuthenticated, setIsAuthenticated, logout }}
        >
            {children}
        </AuthContext.Provider>
    );
};

export default AuthProvider;
