import { ERROR_MESSAGES } from '../../constants';

export const phoneRegExp =
    /^((\+\d{1,3}(-| )?\(?\d\)?(-| )?\d{1,3})|(\(?\d{2,3}\)?))(-| )?(\d{3,4})(-| )?(\d{4})(( x| ext)\d{1,5}){0,1}$/;
export const emailRegExp = /^[\w-.]+@([\w-]+\.)+[\w-]{2,4}$/g;
export const dniRegExp = /^\d{8}$/;
export const cardRegExp = /^\d{16}$/;
export const aliasRegExp = /[a-z]+\.[a-z]+\.[a-z]+$/ig;
export const moneyRegExp = /^([0]{1}\.{1}[0-9]+|[1-9]{1}[0-9]*\.{1}[0-9]+|[0-9]+|0)$/g;

export const validExpiration = (expiration: string): boolean => {
    const currentYear = new Date().getFullYear();
    const currentMonth = new Date().getMonth() + 1;
    const millenium = currentYear.toString().slice(0, 2);
    const month = parseInt(expiration.slice(0, 2), 10);
    const year = parseInt(millenium + expiration.slice(2, 4), 10);

    return (
        Number(year) > currentYear ||
        (Number(year) === currentYear && Number(month) >= currentMonth)
    );
};

export const transformExpiration = (expiration: number): string => {
    const expiryString = expiration.toString();
    const expiryMonth = expiryString.slice(0, 2);
    const expiryYear = expiryString.slice(2, 4);
    const date = new Date();
    date.setFullYear(+expiryYear, +expiryMonth - 1, 1);
    return date.toISOString();
};

export const isValueEmpty = (values: any): boolean =>
    Object.values(values).some((value) => value === '');
export const valuesHaveErrors = (errors: any): boolean =>
    Object.keys(errors).length !== 0;

export const emailValidationConfig = {
    required: {
        value: true,
        message: ERROR_MESSAGES.REQUIRED_FIELD,
    },
    pattern: {
        value: emailRegExp,
        message: ERROR_MESSAGES.INVALID_EMAIL,
    },
};

export const passwordValidationConfig = {
    required: {
        value: true,
        message: ERROR_MESSAGES.REQUIRED_FIELD,
    },
    pattern: {
        value: /^(?=\w*\d)(?=\w*[a-z])\S{6,20}$/,
        message: ERROR_MESSAGES.INVALID_PASSWORD,
    },
    minLength: {
        value: 6,
        message: ERROR_MESSAGES.MIN_LENGTH,
    },
    maxLength: {
        value: 20,
        message: ERROR_MESSAGES.MAX_LENGTH,
    },
};

export const nameValidationConfig = {
    required: {
        value: true,
        message: ERROR_MESSAGES.REQUIRED_FIELD,
    },
    pattern: {
        value: /^[a-zA-Z]+(([',. -][a-zA-Z ])?[a-zA-Z]*)*$/,
        message: ERROR_MESSAGES.INVALID_NAME,
    },
    minLength: {
        value: 2,
        message: ERROR_MESSAGES.MIN_LENGHT_NAME,
    },
    maxLength: {
        value: 20,
        message: ERROR_MESSAGES.MAX_LENGTH,
    },
};

export const phoneValidationConfig = {
    required: {
        value: true,
        message: ERROR_MESSAGES.REQUIRED_FIELD,
    },
    pattern: {
        value: phoneRegExp,
        message: ERROR_MESSAGES.INVALID_PHONE,
    },
};

export const dniValidationConfig = {
    required: {
        value: true,
        message: ERROR_MESSAGES.REQUIRED_FIELD,
    },
    pattern: {
        value: dniRegExp,
        message: ERROR_MESSAGES.INVALID_DNI,
    },
};

export const cardValidationConfig = {
    required: {
        value: true,
        message: ERROR_MESSAGES.REQUIRED_FIELD,
    },
    pattern: {
        value: cardRegExp,
        message: ERROR_MESSAGES.INVALID_CARD,
    },
};

export const expirationValidationConfig = {
    required: {
        value: true,
        message: ERROR_MESSAGES.REQUIRED_FIELD,
    },
    validate: (value: any) => {
        if (!validExpiration(value)) {
            return ERROR_MESSAGES.INVALID_EXPIRATION;
        }
    },
};

export const cvcValidationConfig = {
    required: {
        value: true,
        message: ERROR_MESSAGES.REQUIRED_FIELD,
    },
    minLength: {
        value: 3,
        message: ERROR_MESSAGES.INVALID_CVC,
    },
};

export const aliasValidationConfig = {
    required: {
        value: true,
        message: ERROR_MESSAGES.REQUIRED_FIELD,
    },
    pattern: {
        value: aliasRegExp,
        message: ERROR_MESSAGES.INVALID_ALIAS,
    },
    minLength: {
        value: 10,
        message: ERROR_MESSAGES.MIN_LENGHT_ALIAS,
    },
};

export const moneyValidationConfig = {
    required: {
        value: true,
        message: ERROR_MESSAGES.REQUIRED_FIELD,
    },
    pattern: {
        value: moneyRegExp,
        message: ERROR_MESSAGES.INVALID_MONEY,
    },
    validate: (value: any) => {
        if (value === '0') {
            return ERROR_MESSAGES.INVALID_EMPTY_MONEY;
        }
    },
};

export const handleChange = <T>(
    event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>,
    setValue: React.Dispatch<React.SetStateAction<T>>,
    maxLength?: number
) => {
    const { name, value } = event.target;
    const newValue = maxLength ? value.slice(0, maxLength) : value;
    setValue((prevValue: any) => ({
        ...prevValue,
        [name]: newValue,
    }));
};
