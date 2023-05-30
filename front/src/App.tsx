import React, { Suspense } from 'react';
import { BrowserRouter, Navigate, Route, Routes } from 'react-router-dom';
import { PrivateRoutes } from './components';
import { ROUTES } from './constants';
import { Layout } from './components/Layout';
import './tailwind/styles.css';
import CircularProgress from '@mui/material/CircularProgress';
import Dashboard from './pages/Dashboard';
import { useAuth } from './hooks';

// pages
const Login = React.lazy(() => import('./pages/Login'));
const Register = React.lazy(() => import('./pages/Register'));
const Activity = React.lazy(() => import('./pages/Activity'));
const ActivityDetails = React.lazy(() => import('./pages/ActivityDetails'));
const Cards = React.lazy(() => import('./pages/Cards'));
const SendMoney = React.lazy(() => import('./pages/SendMoney'));
const LoadMoney = React.lazy(() => import('./pages/LoadMoney'));
const Profile = React.lazy(() => import('./pages/Profile'));
const PageNotFound = React.lazy(() => import('./pages/PageNotFound'));

function App() {
    const { isAuthenticated } = useAuth();

    return (
        <>
            <BrowserRouter>
                <Layout isAuthenticated={isAuthenticated}>
                    <Suspense
                        fallback={
                            <div className="tw-w-full tw-h-full tw-flex tw-flex-col tw-items-center tw-justify-center">
                                <CircularProgress />
                            </div>
                        }
                    >
                        <Routes>
                            <React.Fragment></React.Fragment>
                            <Route path={ROUTES.HOME} element={<PrivateRoutes />}>
                                <Route element={<Dashboard />} path={ROUTES.HOME} />
                                <Route element={<Activity />} path={`${ROUTES.ACTIVITY}`} />
                                <Route element={<Cards />} path={ROUTES.CARDS} />
                                <Route element={<SendMoney />} path={ROUTES.SEND_MONEY} />
                                <Route element={<LoadMoney />} path={ROUTES.LOAD_MONEY} />
                                <Route element={<Profile />} path={ROUTES.PROFILE} />
                                <Route
                                    element={<ActivityDetails />}
                                    path={ROUTES.ACTIVITY_DETAILS}
                                />
                            </Route>
                            <Route
                                element={
                                    isAuthenticated ? (
                                        <Navigate replace to={ROUTES.HOME} />
                                    ) : (
                                        <Login />
                                    )
                                }
                                path={ROUTES.LOGIN}
                            />
                            <Route
                                element={
                                    isAuthenticated ? (
                                        <Navigate replace to={ROUTES.HOME} />
                                    ) : (
                                        <Register />
                                    )
                                }
                                path={ROUTES.REGISTER}
                            />
                            <Route element={<PageNotFound />} path="*" />
                        </Routes>
                    </Suspense>
                </Layout>
            </BrowserRouter>
        </>
    );
}

export default App;
