import React, { useEffect } from 'react';
import { Navigate, Outlet } from 'react-router-dom';
import { useLocalStorage } from '../../hooks';
import { useAuth } from '../../hooks/useAuth';

export const PrivateRoutes = () => {
    const [token] = useLocalStorage('token');
    const { isAuthenticated, setIsAuthenticated } = useAuth();
    useEffect(() => {
        if (isAuthenticated) {
            if (token) {
                setIsAuthenticated(true);
            } else {
                setIsAuthenticated(false);
            }
        }
    }, [isAuthenticated, setIsAuthenticated, token]);

    return isAuthenticated ? <Outlet /> : <Navigate to="/login" />;
};
