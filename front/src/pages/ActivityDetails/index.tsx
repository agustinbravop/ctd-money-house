import React, { useEffect, useState } from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { CardCustom, Icon } from '../../components/';
import { currencies, RECORD_MESSAGES, ROUTES, UNAUTHORIZED, } from '../../constants';
import { Button } from '@mui/material';
import {
    calculateTransacionType,
    formatCurrency,
    formatDateFromString,
    getUserActivity,
    printPage,
} from '../../utils';
import { ActivityType, Transaction } from '../../types';
import { useAuth, useLocalStorage, useUserInfo } from '../../hooks';

const ActivityDetails = () => {
    const [userActivity, setUserActivity] = useState<Transaction>();
    const [activityType, setActivityType] = useState<ActivityType>(
        ActivityType.DEPOSIT
    );

    const [searchParams] = useSearchParams();
    const activityId = searchParams.get('id') || '1';
    const navigate = useNavigate();
    const { Argentina } = currencies;
    const { locales, currency } = Argentina;
    const [token] = useLocalStorage('token');
    const { logout } = useAuth();
    const { user } = useUserInfo();

    useEffect(() => {
        if (user && user.id) {
            getUserActivity(user.id, activityId, token)
                .then((activity) => {
                    if (activity && activity.amount && activity.transaction_type) {
                        setUserActivity(activity);
                        setActivityType(
                            calculateTransacionType(activity.amount, activity.transaction_type)
                        );
                    }
                })
                .catch((error) => {
                    if (error.status === UNAUTHORIZED) {
                        logout();
                    }
                });
        }
    }, [activityId, logout, token, user]);

    return (
        <div>
            <CardCustom
                className="tw-flex tw-flex-col tw-justify-center tw-items-center tw-max-w-5xl print:tw-border-none print:tw-w-screen print:tw-h-screen print:!tw-mt-0"
                content={
                    <div className="tw-border-neutral-blue-100 tw-rounded-lg print:tw-border-2 print:tw-p-48">
                        <Icon
                            className="tw-text-primary tw-mx-auto tw-mb-8 tw-hidden print:tw-block"
                            type="digital-house"
                        />
                        <div className="tw-flex tw-flex-col tw-gap-y-6 tw-justify-center tw-mb-8 print:tw-mb-0">
                            <div className="tw-flex tw-flex-col tw-gap-y-2 tw-items-center">
                                <p>Monto</p>
                                <p className="tw-text-xl tw-font-bold">
                                    {userActivity &&
                                        formatCurrency(
                                            locales,
                                            currency,
                                            Math.abs(userActivity.amount)
                                        )}
                                </p>
                            </div>
                            <div className="tw-flex tw-flex-col tw-gap-y-2 tw-items-center">
                                <p>
                                    {RECORD_MESSAGES[activityType] &&
                                        RECORD_MESSAGES[activityType]}{' '}
                                </p>
                                <p className="tw-text-xl tw-font-bold">
                                    {userActivity && userActivity.description && userActivity.description}
                                </p>
                                <p className="tw-text-xl tw-font-bold">
                                    {userActivity &&
                                        userActivity.destination_cvu &&
                                        userActivity.destination_cvu}
                                </p>
                                <i>
                                    {userActivity && formatDateFromString(userActivity.transaction_date)}
                                </i>
                            </div>
                        </div>
                    </div>
                }
                actions={
                    <div className="tw-flex tw-flex-col tw-justify-center tw-gap-y-8">
                        <Button
                            onClick={printPage}
                            className="tw-h-12 tw-w-64 print:tw-hidden"
                            variant="outlined"
                        >
                            Imprimir
                        </Button>
                        <Button
                            onClick={() => navigate(ROUTES.HOME)}
                            className="tw-h-12 tw-w-64 print:tw-hidden"
                            variant="contained"
                        >
                            Continuar
                        </Button>
                    </div>
                }
            />
        </div>
    );
};

export default ActivityDetails;
