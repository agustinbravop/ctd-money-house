import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import UserInfoProvider from './context/User';
import AuthProvider from './context/Auth';

const darkTheme = createTheme({
    palette: {
        mode: 'dark',
    },
});

const root = ReactDOM.createRoot(
    document.getElementById('root') as HTMLElement
);
root.render(
    <React.StrictMode>
        <AuthProvider>
            <UserInfoProvider>
                <ThemeProvider theme={darkTheme}>
                    <App />
                </ThemeProvider>
            </UserInfoProvider>
        </AuthProvider>
    </React.StrictMode>
);
