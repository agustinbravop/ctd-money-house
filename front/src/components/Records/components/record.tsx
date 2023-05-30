import React from 'react';
import {
    calculateTransacionType,
    deleteUserCard,
    formatCurrency,
    formatDateFromString,
    isMastercard,
    isVisa,
} from '../../../utils/';
import {
    currencies,
    DESTINATION,
    ID,
    MESSAGE,
    RECORD_MESSAGES,
    ROUTES,
    STEP,
    SUCCESS,
    SUCCESS_MESSAGES_KEYS,
    UNAUTHORIZED,
} from '../../../constants/';
import { Link, useNavigate, useSearchParams } from 'react-router-dom';
import { Icon, IconType } from '../../Icon';
import { Account, ActivityType, Card, Transaction } from '../../../types';
import { useAuth, useLocalStorage, useUserInfo } from '../../../hooks';

export enum RecordVariant {
    TRANSACTION = 'transaction',
    CARD = 'card',
    ACCOUNT = 'account',
}

export interface IRecord {
    content: Transaction | Card | Account;
    variant?: RecordVariant;
}

export interface RecordProps extends IRecord {
    className?: string;
    setRecords?: React.Dispatch<React.SetStateAction<IRecord[]>>;
}

const { Argentina } = currencies;
const { locales, currency } = Argentina;

const iconType: Record<ActivityType, IconType> = {
    [ActivityType.TRANSFER_IN]: 'transfer-in',
    [ActivityType.TRANSFER_OUT]: 'transfer-out',
    [ActivityType.DEPOSIT]: 'deposit',
};

export const Record = ({
                           className,
                           content,
                           variant = RecordVariant.TRANSACTION,
                           setRecords,
                       }: RecordProps) => {
    const [searchParams] = useSearchParams();
    const isSelecting = !!searchParams.get('select');

    return (
        <li
            className={`tw-flex tw-w-full tw-justify-between tw-px-4 tw-border-t tw-border-neutral-blue-100 tw-py-5 hover:tw-bg-neutral-gray-500 tw-transition ${className}`}
        >
            {variant === RecordVariant.TRANSACTION && (
                <TransactionItem {...(content as Transaction)} />
            )}
            {variant === RecordVariant.CARD && (
                <CardItem
                    {...(content as Card)}
                    setRecords={setRecords}
                    isSelecting={isSelecting}
                />
            )}
            {variant === RecordVariant.ACCOUNT && (
                <AccountItem {...(content as Account)} />
            )}
        </li>
    );
};

function TransactionItem({ amount, description, transaction_date, id, transaction_type }: Transaction) {
    const calculatedType = calculateTransacionType(amount, transaction_type);
    return (
        <Link
            className="tw-w-full tw-flex tw-justify-between tw-items-center"
            to={`${ROUTES.ACTIVITY_DETAILS}?${ID}${id}`}
        >
            <div className="tw-flex tw-items-center tw-gap-x-4">
                {calculatedType && (
                    <Icon
                        className="tw-fill-neutral-gray-500"
                        type={iconType[calculatedType]}
                    />
                )}
                <p>
                    {RECORD_MESSAGES[calculatedType] && RECORD_MESSAGES[calculatedType]}{' '}
                    {description}
                </p>
            </div>
            <div className="tw-flex tw-text-left tw-flex-col tw-items-end">
                <p>{formatCurrency(locales, currency, amount)}</p>
                <p>{formatDateFromString(transaction_date)}</p>
            </div>
        </Link>
    );
}

function CardItem({
                      card_number,
                      type,
                      isSelecting,
                      id: cardId,
                      setRecords,
                  }: Card & {
    setRecords?: React.Dispatch<React.SetStateAction<IRecord[]>>;
    isSelecting: boolean;
}) {
    const navigate = useNavigate();
    const lastFourDigits = (card_number && card_number.slice(-4)) || '';
    const isVisaCard = isVisa(card_number);
    const isMasterCard = isMastercard(card_number);
    const cardType = isVisaCard
        ? 'visa'
        : isMasterCard
            ? 'mastercard'
            : 'credit-card';
    const { user } = useUserInfo();
    const { logout } = useAuth();
    const [token] = useLocalStorage('token');

    const handleDelete = () => {
        if (user && user.id) {
            deleteUserCard(user.id, cardId, token)
                .then((response) => {
                    if (response.status === UNAUTHORIZED) {
                        logout();
                    }
                    if (setRecords) {
                        setRecords((prev) =>
                            prev.filter((record) => (record.content as Card).id !== cardId)
                        );
                    }
                    navigate(
                        `${ROUTES.CARDS}?${SUCCESS}&${MESSAGE}${SUCCESS_MESSAGES_KEYS.CARD_DELETED}`
                    );
                })
                .catch((error) => {
                    if (error.status === UNAUTHORIZED) {
                        logout();
                    }
                });
        }
    };

    return (
        <>
            <div className="tw-flex tw-items-center tw-gap-x-4">
                <Icon type={cardType} />

                <p>
                    {type} terminada en {lastFourDigits}
                </p>
            </div>
            <div className="tw-flex tw-text-left tw-gap-x-4 tw-items-center">
                {isSelecting ? (
                    <button
                        onClick={() =>
                            navigate(
                                `${ROUTES.LOAD_MONEY}?type=${cardType}&card=${cardId}`
                            )
                        }
                        className="tw-text-primary"
                    >
                        Seleccionar
                    </button>
                ) : (
                    <button className="tw-text-error" onClick={handleDelete}>Eliminar</button>
                )}
            </div>
        </>
    );
}

function AccountItem({ name, origin }: Account) {
    const navigate = useNavigate();

    return (
        <>
            <div className="tw-flex tw-items-center tw-gap-x-4">
                <Icon type="user" />
                <p>{name}</p>
            </div>
            <div className="tw-flex tw-text-primary tw-text-left tw-gap-x-4 tw-items-center">
                <button
                    onClick={() =>
                        navigate(`${ROUTES.SEND_MONEY}?${STEP}2&${DESTINATION}${origin}`)
                    }
                >
                    Seleccionar
                </button>
            </div>
        </>
    );
}
