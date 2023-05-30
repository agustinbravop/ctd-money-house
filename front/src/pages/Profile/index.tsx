import React, { useEffect, useMemo, useState } from 'react';
import { Link, useNavigate, useSearchParams } from 'react-router-dom';
import { Button } from '@mui/material';
import FormControl from '@mui/material/FormControl';
import InputLabel from '@mui/material/InputLabel';
import OutlinedInput from '@mui/material/OutlinedInput';
import { SubmitHandler, useForm } from 'react-hook-form';
import {
    EDIT,
    MESSAGE,
    ROUTES,
    SUCCESS,
    SUCCESS_MESSAGES,
    SUCCESS_MESSAGES_KEYS,
    UNAUTHORIZED,
} from '../../constants';
import { CardCustom, ErrorMessage, Errors, Icon, SnackBar, Tooltip, TooltipPosition, } from '../../components';
import {
    aliasValidationConfig,
    copyToClipboard,
    getAccount,
    isValueEmpty,
    updateAccount,
    valuesHaveErrors,
} from '../../utils';
import { useAuth, useLocalStorage, useUserInfo } from '../../hooks';

export interface IProfile {
    alias?: string;
}

const duration = 2000;
const Profile = () => {
    const navigate = useNavigate();
    const [searchParams] = useSearchParams();
    const isEditing = !!searchParams.get('edit');
    const isSuccess = !!searchParams.get('success');
    const message = (searchParams.get('message') as SUCCESS_MESSAGES_KEYS) || '';
    const [isError, setIsError] = useState<boolean>(!!searchParams.get('error'));
    const { user } = useUserInfo();
    const [token, setToken] = useLocalStorage('token');

    const [userAccount, setUserAccount] = useState({
        alias: '',
        cvu: '',
    });

    const {
        register,
        handleSubmit,
        formState: { errors, isDirty },
    } = useForm({
        criteriaMode: 'all',
    });

    const isEmpty = isValueEmpty(userAccount.alias);
    const hasErrors = useMemo(() => valuesHaveErrors(errors), [errors]);
    const { setIsAuthenticated } = useAuth();

    useEffect(() => {
        if (user && user.id) {
            getAccount(token)
                .then((account) => {
                    if (account && account.alias && account.cvu) {
                        setUserAccount(account);
                    }
                    if (isSuccess) {
                        setTimeout(() => navigate(ROUTES.PROFILE), duration);
                    }
                })
                .catch((error) => {
                    if ((error as Response).status === UNAUTHORIZED) {
                        setToken(null);
                    }
                });
        } else {
            setIsAuthenticated(false);
        }
    }, [isSuccess, navigate, setIsAuthenticated, setToken, token, user]);

    const onChange = (
        event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>
    ) => setUserAccount({ ...userAccount, alias: event.target.value });

    const onSubmit: SubmitHandler<IProfile> = (data) => {
        if (user && user.id) {
            updateAccount(user.id, { alias: data.alias }, token)
                .then((response) => {
                    if (response.status) {
                        setIsError(true);
                    } else {
                        navigate(
                            `${ROUTES.PROFILE}?${SUCCESS}&${MESSAGE}${SUCCESS_MESSAGES_KEYS.ALIAS_EDITED}`
                        );
                    }
                })
                .catch((error) => {
                    if ((error as Response).status === UNAUTHORIZED) {
                        setToken(null);
                    }
                });
        }
    };

    return (
        <div className="tw-w-full">
            {!isEditing ? (
                <CardCustom
                    className="tw-max-w-5xl"
                    content={
                        <div className="tw-flex tw-gap-4 tw-flex-col tw-w-full">
                            <p className="tw-font-bold">
                                Copia tu cvu o alias para ingresar o transferir dinero desde
                                otra cuenta
                            </p>
                            <div className="tw-flex tw-mb-4 tw-justify-between tw-items-center">
                                <div>
                                    <p className="tw-font-bold tw-text-primary">CVU</p>
                                    <p className="">{userAccount.cvu}</p>
                                </div>
                                <Tooltip
                                    className="tw-cursor-pointer"
                                    message="Copiado"
                                    position={TooltipPosition.top}
                                >
                                    <button
                                        onClick={() => copyToClipboard(userAccount.cvu || '')}
                                    >
                                        <Icon type="copy" />
                                    </button>
                                </Tooltip>
                            </div>
                            <div className="tw-flex tw-justify-between tw-items-center">
                                <div>
                                    <p className="tw-font-bold tw-text-primary">Alias</p>
                                    <p className="">{userAccount.alias}</p>
                                </div>
                                <div className="tw-flex">
                                    <Link to={`${ROUTES.PROFILE}?${EDIT}`}>
                                        <Icon className="tw-text-primary" type="edit" />
                                    </Link>
                                    <Tooltip
                                        className="tw-cursor-pointer tw-ml-4"
                                        message="Copiado"
                                        position={TooltipPosition.top}
                                    >
                                        <button
                                            onClick={() => copyToClipboard('estealias.no.existe')}
                                        >
                                            <Icon type="copy" />
                                        </button>
                                    </Tooltip>
                                </div>
                            </div>
                        </div>
                    }
                />
            ) : (
                <CardCustom
                    className="tw-max-w-5xl"
                    content={
                        <div className="tw-flex tw-flex-col">
                            <p className="tw-font-bold tw-mb-4">Editar alias</p>

                            <form onSubmit={handleSubmit(onSubmit)}>
                                <div>
                                    <FormControl variant="outlined">
                                        <InputLabel htmlFor="outlined-adornment-alias">
                                            Alias
                                        </InputLabel>
                                        <OutlinedInput
                                            id="outlined-adornment-alias"
                                            type="text"
                                            value={userAccount.alias}
                                            {...register('alias', aliasValidationConfig)}
                                            onChange={onChange}
                                            label="alias"
                                            autoComplete="off"
                                        />
                                    </FormControl>
                                    {errors.alias && (
                                        <ErrorMessage errors={errors.alias as Errors} />
                                    )}
                                    <div className="tw-flex tw-w-full tw-justify-end tw-mt-6">
                                        <Button
                                            type="submit"
                                            className={`tw-h-12 tw-w-64 ${
                                                hasErrors || !isDirty || isEmpty
                                                    ? 'tw-text-neutral-gray-300 tw-border-neutral-gray-300 tw-cursor-not-allowed'
                                                    : ''
                                            }`}
                                            variant="outlined"
                                            disabled={hasErrors || !isDirty || isEmpty}
                                        >
                                            Confirmar
                                        </Button>
                                    </div>
                                </div>
                            </form>
                        </div>
                    }
                />
            )}
            {isSuccess && (
                <SnackBar
                    duration={duration}
                    message={SUCCESS_MESSAGES[message] ? SUCCESS_MESSAGES[message] : ''}
                    type="success"
                />
            )}
            {isError && (
                <SnackBar
                    duration={duration}
                    message="El alias seleccionado ya existe. Debe ingresar uno nuevo"
                    type="error"
                />
            )}
        </div>
    );
};

export default Profile;
