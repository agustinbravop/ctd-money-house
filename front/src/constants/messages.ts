import { ActivityType } from '../types';

export const ERROR_MESSAGES: Record<string, string> = {
    INVALID_EMAIL: 'Correo electrónico inválido',
    INVALID_PASSWORD: 'Contraseña inválida',
    PASSWORDS_DO_NOT_MATCH: 'Las contraseñas no coinciden',
    INVALID_NAME: 'Nombre inválido',
    INVALID_PHONE: 'Teléfono inválido',
    INVALID_DNI: 'DNI inválido',
    INVALID_CARD: 'Tarjeta inválida',
    INVALID_EXPIRATION: 'Fecha de expiración inválida',
    INVALID_CVC: 'CVC debe tener al menos 3 dígitos',
    INVALID_ALIAS: 'El alias deben ser 3 palabras separadas por puntos',
    INVALID_MONEY: 'El monto a ingresar no puede ser negativo',
    INVALID_EMPTY_MONEY: 'El monto a ingresar no puede ser cero',
    REQUIRED_FIELD: 'Campo requerido',
    MIN_LENGTH: 'Debe tener al menos 6 caracteres',
    MAX_LENGTH: 'Debe tener menos de 20 caracteres',
    MIN_LENGHT_NAME: 'Debe tener al menos 2 caracteres',
    MIN_LENGHT_ALIAS: 'Debe tener al menos 10 caracteres',
    MIN_LENGHT_MONEY: 'El mínimo para ingresar a la cuenta es de $100',
    INVALID_USER: 'El usuario ya existe',
    NOT_FOUND_USER: 'Usuario no encontrado',
};

export enum SUCCESS_MESSAGES_KEYS {
    CARD_DELETED = 'CARD_DELETED',
    ALIAS_EDITED = 'ALIAS_EDITED',
    CARD_ADDED = 'CARD_ADDED',
    USER_REGISTER = 'USER_REGISTER',
}

export const SUCCESS_MESSAGES: Record<SUCCESS_MESSAGES_KEYS, string> = {
    CARD_DELETED: 'Tarjeta eliminada correctamente',
    ALIAS_EDITED: 'El alias se actualizó correctamente',
    CARD_ADDED: 'Tarjeta agregada correctamente',
    USER_REGISTER: 'Usuario registrado correctamente',
};

export const RECORD_MESSAGES: Record<ActivityType, string> = {
    [ActivityType.TRANSFER_IN]: 'Recibiste de',
    [ActivityType.TRANSFER_OUT]: 'Enviaste dinero',
    [ActivityType.DEPOSIT]: 'Ingresaste',
};
