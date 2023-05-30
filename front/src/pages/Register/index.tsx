import React, { useMemo, useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import Button from '@mui/material/Button';
import FormControl from '@mui/material/FormControl';
import InputLabel from '@mui/material/InputLabel';
import OutlinedInput from '@mui/material/OutlinedInput';
import InputAdornment from '@mui/material/InputAdornment';
import IconButton from '@mui/material/IconButton';
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';
import {
    createAnUser,
    dniValidationConfig,
    emailValidationConfig,
    handleChange,
    isValueEmpty,
    nameValidationConfig,
    passwordValidationConfig,
    phoneValidationConfig,
    valuesHaveErrors,
} from '../../utils/';
import { ErrorMessage, Errors } from '../../components/ErrorMessage';
import { SnackBar } from '../../components';
import { BAD_REQUEST, ERROR_MESSAGES, ROUTES, SUCCESS_MESSAGES, SUCCESS_MESSAGES_KEYS, } from '../../constants/';
import { useLocalStorage } from '../../hooks';
import { useNavigate } from 'react-router-dom';

interface RegisterState {
    name: string;
    lastName: string;
    phone: string;
    dni: string;
    email: string;
    password: string;
    passwordRepeated: string;
    showPassword: boolean;
}

interface RegisterInputs {
    name: string;
    lastName: string;
    phone: string;
    dni: string;
    email: string;
    password: string;
    passwordRepeated: string;
}

const messageDuration = 2000;

const Register = () => {
    const {
        register,
        handleSubmit,
        watch,
        formState: { errors, isDirty },
    } = useForm<RegisterInputs>({
        criteriaMode: 'all',
    });

    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    const [token, setToken] = useLocalStorage('token');
    // const { setIsAuthenticated } = useAuth();

    const [values, setValues] = React.useState<RegisterState>({
        email: '',
        password: '',
        name: '',
        lastName: '',
        phone: '',
        dni: '',
        passwordRepeated: '',
        showPassword: false,
    });
    const [isSuccess, setIsSuccess] = useState<boolean>(false);
    const [isError, setIsError] = useState<boolean>(false);
    const [isSubmiting, setIsSubmiting] = useState<boolean>(false);
    const [message, setMessage] = useState<string>('');

    const isEmpty = isValueEmpty(values);
    const hasErrors = useMemo(() => valuesHaveErrors(errors), [errors]);
    const navigate = useNavigate();

    const handleClickShowPassword = () => {
        setValues({
            ...values,
            showPassword: !values.showPassword,
        });
    };

    const handleMouseDownPassword = (
        event: React.MouseEvent<HTMLButtonElement>
    ) => {
        event.preventDefault();
    };

    const onChange = (
        event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
        maxLength?: number
    ) => handleChange<RegisterState>(event, setValues, maxLength);

    const onSubmit: SubmitHandler<RegisterInputs> = ({
                                                         name,
                                                         lastName,
                                                         password,
                                                         phone,
                                                         dni,
                                                         email,
                                                     }) => {
        setIsSubmiting(true);
        // API: input del dni es numérico, la conversión no hace falta. La comentamos.
        // const parsedDni = parseInt(String(dni));

        createAnUser({
            firstName: name,
            lastName,
            password,
            telephone: phone,
            dni,
            email,
        })
            .then((_) => {
                setIsSuccess(true);
                // API: /auth/register no devuelve un access_token. Comentamos la linea.
                // setToken(response.accessToken);
                setMessage(SUCCESS_MESSAGES[SUCCESS_MESSAGES_KEYS.USER_REGISTER]);
                setTimeout(() => {
                    setIsSubmiting(false);
                    // setIsAuthenticated(true);
                    navigate(ROUTES.LOGIN);
                }, messageDuration);
            })
            .catch((error) => {
                setIsError(true);
                setMessage(ERROR_MESSAGES.INVALID_USER);
                setIsSubmiting(false);
                if (error.status === BAD_REQUEST) {
                    setIsError(true);
                }
            });
    };

    return (
        <div className="tw-w-full tw-h-full tw-flex tw-flex-col tw-flex-1 tw-items-center tw-justify-center">
            <h2>Crear cuenta</h2>
            <div className="tw-flex tw-max-w-3xl">
                <form
                    className="tw-flex tw-flex-wrap tw-gap-x-16 tw-gap-y-12 tw-mt-10 tw-bg-background tw-justify-between"
                    onSubmit={handleSubmit(onSubmit)}
                >
                    <div>
                        <FormControl variant="outlined">
                            <InputLabel htmlFor="outlined-adornment-password">
                                Nombre
                            </InputLabel>
                            <OutlinedInput
                                id="outlined-adornment-name"
                                type="text"
                                value={values.name}
                                {...register('name', nameValidationConfig)}
                                onChange={onChange}
                                label="nombre"
                            />
                        </FormControl>
                        {errors.name && <ErrorMessage errors={errors.name as Errors} />}
                    </div>
                    <div>
                        <FormControl variant="outlined">
                            <InputLabel htmlFor="outlined-adornment-password">
                                Apellido
                            </InputLabel>
                            <OutlinedInput
                                id="outlined-adornment-last-name"
                                type="text"
                                value={values.lastName}
                                {...register('lastName', nameValidationConfig)}
                                onChange={onChange}
                                label="lastName"
                            />
                        </FormControl>
                        {errors.lastName && (
                            <ErrorMessage errors={errors.lastName as Errors} />
                        )}
                    </div>
                    <div>
                        <FormControl variant="outlined">
                            <InputLabel htmlFor="outlined-adornment-dni">DNI</InputLabel>
                            <OutlinedInput
                                id="outlined-adornment-dni"
                                type="number"
                                value={values.dni}
                                {...register('dni', dniValidationConfig)}
                                onChange={(event) => onChange(event, 8)}
                                label="dni"
                                autoComplete="off"
                            />
                        </FormControl>
                        {errors.dni && <ErrorMessage errors={errors.dni as Errors} />}
                    </div>

                    <div>
                        <FormControl variant="outlined">
                            <InputLabel htmlFor="outlined-adornment-password">
                                Correo
                            </InputLabel>
                            <OutlinedInput
                                id="outlined-adornment-email"
                                type="text"
                                value={values.email}
                                {...register('email', emailValidationConfig)}
                                onChange={onChange}
                                label="email"
                            />
                        </FormControl>
                        {errors.email && <ErrorMessage errors={errors.email as Errors} />}
                    </div>
                    <div>
                        <FormControl variant="outlined">
                            <InputLabel htmlFor="outlined-adornment-password">
                                Contraseña
                            </InputLabel>
                            <OutlinedInput
                                id="outlined-adornment-password"
                                type={values.showPassword ? 'text' : 'password'}
                                value={values.password}
                                {...register('password', passwordValidationConfig)}
                                onChange={onChange}
                                endAdornment={
                                    <InputAdornment position="end">
                                        <IconButton
                                            aria-label="toggle password visibility"
                                            onClick={handleClickShowPassword}
                                            onMouseDown={handleMouseDownPassword}
                                            edge="end"
                                            className="tw-text-neutral-gray-100"
                                        >
                                            {values.showPassword ? <VisibilityOff /> : <Visibility />}
                                        </IconButton>
                                    </InputAdornment>
                                }
                                label="Password"
                                autoComplete="off"
                            />
                        </FormControl>
                        {errors.password && (
                            <ErrorMessage errors={errors.password as Errors} />
                        )}
                    </div>

                    <div>
                        <FormControl variant="outlined">
                            <InputLabel htmlFor="outlined-adornment-password-repeated">
                                Confirmar contraseña
                            </InputLabel>
                            <OutlinedInput
                                id="outlined-adornment-password-repeated"
                                type={values.showPassword ? 'text' : 'password'}
                                value={values.passwordRepeated}
                                {...register('passwordRepeated', {
                                    validate: (value: string) => {
                                        if (watch('password') !== value) {
                                            return ERROR_MESSAGES.PASSWORDS_DO_NOT_MATCH;
                                        }
                                    },
                                })}
                                onChange={onChange}
                                endAdornment={
                                    <InputAdornment position="end">
                                        <IconButton
                                            aria-label="toggle password visibility"
                                            onClick={handleClickShowPassword}
                                            onMouseDown={handleMouseDownPassword}
                                            edge="end"
                                            className="tw-text-neutral-gray-100"
                                        >
                                            {values.showPassword ? <VisibilityOff /> : <Visibility />}
                                        </IconButton>
                                    </InputAdornment>
                                }
                                label="Password"
                                autoComplete="off"
                            />
                        </FormControl>
                        {errors.password && (
                            <ErrorMessage errors={errors.passwordRepeated as Errors} />
                        )}
                    </div>

                    <div>
                        <FormControl variant="outlined">
                            <InputLabel htmlFor="outlined-adornment-dni">Télefono</InputLabel>
                            <OutlinedInput
                                id="outlined-adornment-phone"
                                type="number"
                                value={values.phone}
                                {...register('phone', phoneValidationConfig)}
                                onChange={onChange}
                                label="phone"
                            />
                        </FormControl>
                        {errors.phone && <ErrorMessage errors={errors.phone as Errors} />}
                    </div>
                    <div className="tw-w-full tw-flex tw-justify-center">
                        <Button
                            className={`tw-h-14 tw-w-80 ${
                                hasErrors || !isDirty || isEmpty || isSubmiting
                                    ? 'tw-text-neutral-gray-300 tw-border-neutral-gray-300 tw-cursor-not-allowed'
                                    : ''
                            }`}
                            type="submit"
                            variant="outlined"
                            disabled={hasErrors || !isDirty || isEmpty || isSubmiting}
                        >
                            Ingresar
                        </Button>
                    </div>
                </form>
            </div>
            {message.length > 0 && (
                <SnackBar
                    duration={messageDuration}
                    message={message}
                    type={isSuccess ? 'success' : isError ? 'error' : 'primary'}
                />
            )}
        </div>
    );
};

export default Register;
