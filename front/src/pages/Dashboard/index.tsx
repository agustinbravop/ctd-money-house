import React, { useEffect, useState } from 'react';
import Button from '@mui/material/Button';
import { Link, useNavigate, useSearchParams } from 'react-router-dom';
import { ROUTES } from '../../constants';
import {
    CardCustom,
    Icon,
    IRecord,
    Records,
    RecordVariant,
    Skeleton,
    SkeletonVariant,
    SnackBar,
} from '../../components';
import { formatCurrency, getAccount, getUserActivities, parseRecordContent, sortByDate, } from '../../utils/';
import { currencies, UNAUTHORIZED } from '../../constants/';
import { useUserInfo } from '../../hooks/useUserInfo';
import { Transaction, UserAccount } from '../../types/';
import { useAuth, useLocalStorage } from '../../hooks';

const numberOfActivities = 5;
const duration = 2000;

const Dashboard = () => {
    const { Argentina } = currencies;
    const { locales, currency } = Argentina;
    const navigate = useNavigate();
    const { user } = useUserInfo();
    const [token] = useLocalStorage('token');
    const [searchParams] = useSearchParams();
    const isSuccess = !!searchParams.get('success');
    const [userActivities, setUserActivities] = useState<IRecord[]>([]);
    const [userAccount, setUserAccount] = useState({
        amount: 0,
    });
    const [isLoading, setIsLoading] = useState<boolean>(true);
    const { logout } = useAuth();

    useEffect(() => {
        if (user && user.id) {
            getUserActivities(user.id, token)
                .then((activities) => {
                    if ((activities as Transaction[]).length > 0) {
                        const orderedActivities = sortByDate(activities);
                        const parsedRecords = orderedActivities.map(
                            (activity: Transaction) =>
                                parseRecordContent(activity, RecordVariant.TRANSACTION)
                        );
                        setUserActivities(parsedRecords);
                    }
                })
                .finally(() => setIsLoading(false))
                .catch((error) => {
                    if (error.status === UNAUTHORIZED) {
                        logout();
                    }
                });
        }
    }, [logout, token, user]);

    useEffect(() => {
        if ((user && user.id) || (user && user.id && isSuccess)) {
            getAccount(token)
                .then((account) => {
                    if ((account as UserAccount).amount) {
                        setUserAccount({
                            amount: account.amount || 0,
                        });
                    }
                    if (isSuccess) {
                        setTimeout(() => navigate(ROUTES.HOME), duration);
                    }
                })
                .catch((error) => {
                    if (error.status === UNAUTHORIZED) {
                        logout();
                    }
                });
        }
    }, [isSuccess, logout, navigate, token, user]);

    return (
        <div className="tw-w-full">
            <CardCustom
                content={
                    <div className="tw-flex tw-justify-between tw-mb-8">
                        <div>
                            <p className="tw-mb-4 tw-font-bold">Dinero disponible</p>
                            <p className="tw-text-xl tw-font-bold">
                                {// API: El autoreformateo de 'balance' a 'amount' se olvidó esta linea
                                    formatCurrency(locales, currency, userAccount?.amount)}
                            </p>
                        </div>
                        <div className="tw-flex tw-justify-between tw-items-start tw-flex-wrap tw-gap-x-4">
                            <Link
                                className="tw-underline hover:tw-text-primary"
                                to={ROUTES.CARDS}
                            >
                                Ver tarjetas
                            </Link>
                            <Link
                                className="tw-underline hover:tw-text-primary"
                                to={ROUTES.PROFILE}
                            >
                                Ver CVU
                            </Link>
                        </div>
                    </div>
                }
                actions={
                    <>
                        <Button
                            onClick={() => navigate(ROUTES.LOAD_MONEY)}
                            className="tw-h-12 tw-w-64"
                            variant="outlined"
                        >
                            Ingresar dinero
                        </Button>
                        <Button
                            onClick={() => navigate(ROUTES.SEND_MONEY)}
                            className="tw-h-12 tw-w-64"
                            variant="contained"
                        >
                            Transferir dinero
                        </Button>
                    </>
                }
            />

            <CardCustom
                content={
                    <>
                        <div>
                            <p className="tw-mb-4 tw-font-bold">Tu actividad reciente</p>
                        </div>
                        {userActivities.length === 0 && !isLoading && (
                            <p>No hay actividad registrada</p>
                        )}
                        {userActivities.length > 0 && (
                            <Records records={userActivities} maxRecords={5} />
                        )}

                        {isLoading && (
                            <Skeleton
                                variant={SkeletonVariant.RECORD_LIST}
                                numberOfItems={numberOfActivities}
                            />
                        )}
                    </>
                }
                actions={
                    userActivities.length === 0 && !isLoading ? null : (
                        <Link
                            to={ROUTES.ACTIVITY}
                            className="tw-h-12 tw-w-full tw-flex tw-items-center tw-justify-between tw-mt-4 hover:tw-text-primary tw-px-4 hover:tw-bg-neutral-gray-500 tw-transition"
                        >
                            <span>Ver toda tu actividad</span>
                            <Icon type="arrow-right" />
                        </Link>
                    )
                }
            />
            {isSuccess && (
                <SnackBar
                    duration={duration}
                    message="El dinero fue ingresado correctamente"
                    type="success"
                />
            )}
        </div>
    );
};

export default Dashboard;
