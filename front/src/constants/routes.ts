export const ROUTES = {
    HOME: '/',
    LOGIN: '/login',
    REGISTER: '/register',
    PROFILE: '/profile',
    ACTIVITY: '/activities',
    ACTIVITY_DETAILS: '/activity',
    LOAD_MONEY: '/load-money',
    SEND_MONEY: '/send-money',
    CARDS: '/cards',
    NOT_FOUND: '/404',
};

export const LINK_LIST = [
    {
        name: 'Inicio',
        href: ROUTES.HOME,
    },
    {
        name: 'Actividad',
        href: ROUTES.ACTIVITY,
    },
    {
        name: 'Tu perfil',
        href: ROUTES.PROFILE,
    },
    {
        name: 'Cargar Dinero',
        href: ROUTES.LOAD_MONEY,
    },
    {
        name: 'Enviar Dinero',
        href: ROUTES.SEND_MONEY,
    },
    {
        name: 'Mis Tarjetas',
        href: ROUTES.CARDS,
    },
    {
        name: 'Cerrar Sesión',
        href: ROUTES.LOGIN,
    }
];