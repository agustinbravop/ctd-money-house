import React, { useCallback, useEffect, useState } from 'react';
import { CardCustom, FormSingle, Icon, IRecord, Records, RecordVariant, SnackBar, } from '../../components';
import { currencies, ERROR, ROUTES, STEP, UNAUTHORIZED } from '../../constants';
import { Link, useNavigate, useSearchParams } from 'react-router-dom';
import {
    calculateTransacionType,
    createTransferActivity,
    formatCurrency,
    getAccount,
    getAccountByAliasOrCvu,
    getUserActivities,
    handleChange,
    moneyValidationConfig,
    parseRecordContent,
} from '../../utils';
import { Button } from '@mui/material';
import { ActivityType, Transaction, TransactionType, UserAccount, } from '../../types';
import { useAuth, useLocalStorage, useUserInfo } from '../../hooks';

const SendMoney = () => {
    const [searchParams] = useSearchParams();
    const step = searchParams.get('step');
    const [userActivities, setUserActivities] = useState<Transaction[]>([]);
    const [userAccounts, setUserAccounts] = useState<IRecord[]>([]);
    const { logout } = useAuth();
    const { user } = useUserInfo();
    const [token] = useLocalStorage('token');

    useEffect(() => {
        if (user && user.id && !step) {
            getUserActivities(user.id, token)
                .then((activities) => {
                    if ((activities as Transaction[]).length > 0) {
                        const parsedActivities = activities.filter(
                            (activity: Transaction) =>
                                activity.transaction_type === TransactionType.Transfer
                        );
                        setUserActivities(parsedActivities);
                    }
                })
                .catch((error) => {
                    if (error.status === UNAUTHORIZED) {
                        logout();
                    }
                });
        }
    }, [logout, step, token, user]);

    useEffect(() => {
        if (userActivities.length > 0) {
            const parsedRecords = userActivities
                .filter(
                    (activity: Transaction) =>
                        calculateTransacionType(activity.amount, activity.transaction_type) ===
                        ActivityType.TRANSFER_IN
                )
                .map((activity: Transaction) =>
                    parseRecordContent(
                        { name: activity.description, origin: activity.origin_cvu },
                        RecordVariant.ACCOUNT
                    )
                );

            const uniqueRecords = parsedRecords.filter(
                (record, index, self) =>
                    index ===
                    self.findIndex(
                        (t) =>
                            (t.content as Transaction).origin_cvu ===
                            (record.content as Transaction).origin_cvu
                    )
            );

            setUserAccounts(uniqueRecords);
        }
    }, [userActivities]);

    return (
        <div className="tw-w-full">
            {step ? (
                <SendMoneyForm />
            ) : (
                <>
                    <CardCustom
                        className="tw-max-w-5xl"
                        content={
                            <div className="tw-flex tw-justify-between tw-mb-4">
                                <p className="tw-font-bold">Elegí a qué cuenta transferir</p>
                            </div>
                        }
                        actions={
                            <Link
                                to={`${ROUTES.SEND_MONEY}?${STEP}1`}
                                className="tw-w-full tw-flex tw-items-center tw-justify-between tw-p-4 hover:tw-bg-neutral-gray-500 tw-transition"
                            >
                                <div className="tw-flex tw-items-center tw-gap-x-4">
                                    <Icon type="add" />
                                    <p>Nueva cuenta</p>
                                </div>
                                <Icon type="arrow-right" />
                            </Link>
                        }
                    />
                    <CardCustom
                        className="tw-max-w-5xl"
                        content={
                            <>
                                <div>
                                    <p className="tw-mb-4 tw-font-bold">Últimas cuentas</p>
                                </div>
                                {userAccounts.length > 0 && <Records records={userAccounts} />}
                                {userAccounts.length === 0 && <p>No hay cuentas registradas</p>}
                            </>
                        }
                    />
                </>
            )}
        </div>
    );
};

export default SendMoney;

const duration = 2000;

function SendMoneyForm() {
    const [searchParams] = useSearchParams();
    const [destination, setDestination] = useState(
        searchParams.get('destination')
    );
    const step = searchParams.get('step');
    const [formState, setFormState] = useState({
        destination: destination || '',
        amount: '',
    });
    const [userDestinationAccount, setUserDestinationAccount] =
        useState<UserAccount>();
    const [userOriginAccount, setUserOriginAccount] = useState<UserAccount>();
    const navigate = useNavigate();
    const { user } = useUserInfo();
    const [token] = useLocalStorage('token');

    useEffect(() => {
        if (destination) {
            // API: Usada la función para buscar por Cvu o Alias. Creada exclusivamente para esto.
            getAccountByAliasOrCvu(destination, token).then((account) => {
                if (account) {
                    setUserDestinationAccount(account);
                } else {
                    setDestination(null);
                    navigate(`${ROUTES.SEND_MONEY}?${STEP}1&${ERROR}`);
                }
            });
        }
    }, [destination, navigate, token]);

    useEffect(() => {
        if (user && user.id) {
            getAccount(token).then((userAccount) => {
                if (userAccount) {
                    setUserOriginAccount(userAccount);
                }
            });
        }
    }, [token, user]);

    const onNavigate = useCallback(
        (step: number) => {
            navigate(`${ROUTES.SEND_MONEY}?${STEP}${step}`);
        },
        [navigate]
    );

    useEffect(() => {
        if (step !== '1' && formState.destination === '') {
            setTimeout(() => onNavigate(1));
        }
    }, [step, formState, navigate, onNavigate]);

    const onChange = (
        event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
        maxLength?: number
    ) => handleChange(event, setFormState, maxLength);

    return <>{renderStep(formState)}</>;

    function renderStep(formState: { destination: string; amount: string }) {
        const { amount } = formState;
        const { Argentina } = currencies;
        const { locales, currency } = Argentina;
        const isError = !!searchParams.get('error');

        const handleClick = (
            origin: string,
            destination: string,
            amount: number
        ) => {
            const description = userDestinationAccount
                ? `Transferencia a ${userDestinationAccount.alias}`
                : '';
            if (user && user.id) {
                createTransferActivity(
                    user.id,
                    token,
                    origin,
                    destination,
                    amount,
                    description
                ).then((response) => {
                    navigate(`${ROUTES.ACTIVITY_DETAILS}?${STEP}${response.id}`);
                });
            }
        };

        switch (step) {
            case '1':
                return (
                    <>
                        <FormSingle
                            name="destination"
                            title="Agregá una nueva cuenta"
                            label="CVU ó Alias"
                            type="text"
                            actionLabel="Continuar"
                            formState={formState}
                            handleChange={(
                                event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
                            ) => onChange(event)}
                            submit={() => {
                                if (formState.destination) {
                                    setDestination(formState.destination);
                                    onNavigate(2);
                                }
                            }}
                        />
                        {isError && (
                            <SnackBar
                                duration={duration}
                                message="Cuenta no encontrada"
                                type="error"
                            />
                        )}
                    </>
                );
            case '2':
                return (
                    <FormSingle
                        name="amount"
                        title={`¿Cuanto quieres transferir a ${
                            userDestinationAccount && userDestinationAccount.cvu
                        } ?`}
                        label="Monto"
                        type="number"
                        actionLabel="Continuar"
                        validation={moneyValidationConfig}
                        formState={formState}
                        handleChange={(
                            event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
                        ) => onChange(event)}
                        submit={() => onNavigate(3)}
                    />
                );
            case '3':
                return (
                    <>
                        <CardCustom
                            content={
                                <div className="tw-flex tw-flex-col">
                                    <div className="tw-flex tw-justify-between tw-mb-4">
                                        <p className="tw-font-bold">Revisá que está todo bien</p>
                                    </div>
                                    <div className="tw-flex tw-flex-col tw-mb-4">
                                        <div className="tw-flex tw-gap-2">
                                            <p className="">Vas a transferir</p>
                                            <Link to={`${ROUTES.SEND_MONEY}?${STEP}2`}>
                                                <Icon type="edit" />
                                            </Link>
                                        </div>
                                        <p className="tw-font-bold">
                                            {formatCurrency(locales, currency, parseFloat(amount))}
                                        </p>
                                    </div>
                                    <div className="tw-flex tw-flex-col tw-mb-4">
                                        <p className="">Para</p>
                                        <p className="tw-font-bold">
                                            {userDestinationAccount &&
                                                userDestinationAccount.cvu}
                                        </p>
                                    </div>
                                    <div className="tw-flex tw-flex-col tw-mb-4">
                                        <p className="tw-font-bold">CVU: {formState.destination}</p>
                                    </div>
                                </div>
                            }
                            actions={
                                <div className="tw-flex tw-w-full tw-justify-end tw-mt-6">
                                    <Button
                                        type="submit"
                                        className="tw-h-12 tw-w-64"
                                        variant="outlined"
                                        onClick={() =>
                                            handleClick(
                                                userOriginAccount?.cvu || '',
                                                userDestinationAccount?.cvu || '',
                                                parseFloat(amount)
                                            )
                                        }
                                    >
                                        Transferir
                                    </Button>
                                </div>
                            }
                        />
                    </>
                );
        }
    }
}
