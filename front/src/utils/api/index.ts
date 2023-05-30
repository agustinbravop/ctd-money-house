import { Card, Transaction, User, UserAccount } from '../../types';

const myInit = (method = 'GET', token?: string) => {
    return {
        method,
        headers: {
            'Content-Type': 'application/json',
            Authorization: token ? `Bearer ${token}` : '',
        },
        mode: 'cors' as RequestMode,
        cache: 'default' as RequestCache,
    };
};

const myRequest = (endpoint: string, method: string, token?: string) =>
    new Request(endpoint, myInit(method, token));

const API_URL = process.env.REACT_APP_API_BASE_URL;

const rejectPromise = (response?: Response): Promise<Response> =>
    Promise.reject({
        status: (response && response.status) || '00',
        statusText: (response && response.statusText) || 'Ocurrió un error',
        err: true,
    });

export const login = (email: string, password: string) => {
    return fetch(myRequest(`${API_URL}/auth/login`, 'POST'), {
        body: JSON.stringify({ email, password }),
    })
        .then((response) => {
            if (response.ok) {
                return response.json();
            }
            return rejectPromise(response);
        })
        .catch((err) => {
            console.log(err);
            return rejectPromise(err);
        });
};

export const createAnUser = (user: User) => {
    return fetch(myRequest(`${API_URL}/auth/register`, 'POST'), {
        body: JSON.stringify(user),
    })
        .then((response) => {
            if (response.ok) {
                return response.json();
            }
            return rejectPromise(response);
        })
        .then((data) => {
            // API: se pasa user (request body) porque tiene password, data (response body) no tiene.
            createAnAccount(user);
            return data;
        })
        .catch((err) => {
            console.log(err);
            return rejectPromise(err);
        });
};

// API: getUser no recibía token por parámetro pero es necesario para el endpoint, asi que se refactorizó eso.
// Busquen los que consumen esta función y agregenle el segundo parámetro (token).
export const getUser = (id: string, token: string): Promise<User> => {
    return fetch(myRequest(`${API_URL}/users/${id}`, 'GET', token))
        .then((response) =>
            response.ok ? response.json() : rejectPromise(response)
        )
        .catch((err) => {
            console.log(err);
            return rejectPromise(err);
        });
};

export const updateUser = (
    id: string,
    data: any,
    token: string
): Promise<Response> => {
    return fetch(myRequest(`${API_URL}/users/${id}`, 'PATCH', token), {
        body: JSON.stringify(data),
    })
        .then((response) =>
            response.ok ? response.json() : rejectPromise(response)
        )
        .catch((err) => {
            console.log(err);
            return rejectPromise(err);
        });
};

// API: createAnAccount se llama cuando se llama a createAnUser.
// Se adaptaron nombres de los campos. Eliminadas funciones generateCvu y generateAlias.
export const createAnAccount = (user: User): Promise<Response> => {
    // API: para crear un account necesita estar logeado (tener un access_token).
    return login(user.email, user.password)
        .then((data) => {
            return data.access_token;
        })
        .then((token) => {
            return fetch(myRequest(`${API_URL}/accounts/`, 'POST', token), {
                body: JSON.stringify({}),
            }).then((response) =>
                response.ok ? response.json() : rejectPromise(response)
            );
        })
        .then((data) => data);
};

export const getAccount = (token: string): Promise<UserAccount> => {
    return fetch(myRequest(`${API_URL}/accounts/`, 'GET', token))
        .then((response) => {
            if (response.ok) {
                // API: parece que acá leía la primer account de la lista de accounts. Nosotros solo devolvemos una.
                return response.json();
            }
            return rejectPromise(response);
        })
        .catch((err) => {
            console.log(err);
            return rejectPromise(err);
        });
};

// API: cambiado getAccounts() por getAccountByCvu(). Esta función es usada en '/src/pages/SendMoney/index.tsx'.
export const getAccountByAliasOrCvu = (
    cvuOrAlias: string,
    token: string
): Promise<UserAccount> => {
    return fetch(
        myRequest(`${API_URL}/accounts/byAliasOrCvu/${cvuOrAlias}`, 'GET', token)
    )
        .then((response) =>
            response.ok ? response.json() : rejectPromise(response)
        )
        .catch((err) => {
            console.log(err);
            return rejectPromise(err);
        });
};

export const updateAccount = (
    id: string,
    data: any,
    token: string
): Promise<Response> => {
    return fetch(myRequest(`${API_URL}/accounts/${id}`, 'PATCH', token), {
        body: JSON.stringify(data),
    })
        .then((response) =>
            response.ok ? response.json() : rejectPromise(response)
        )
        .catch((err) => {
            console.log(err);
            return rejectPromise(err);
        });
};

export const getUserActivities = (
    userId: string,
    token: string,
    limit?: number
): Promise<Transaction[]> => {
    return getAccount(token).then((account) => {
        return fetch(
            myRequest(
                `${API_URL}/accounts/${account.id}/activity${limit ? `?limit=${limit}` : ''}`,
                'GET',
                token
            )
        )
            .then((response) => {
                if (response.ok) {
                    return response.json();
                }
                return rejectPromise(response);
            })
            .catch((err) => {
                console.log(err);
                return rejectPromise(err);
            });
    });
};

export const getUserActivity = (
    userId: string,
    activityId: string,
    token: string
): Promise<Transaction> => {
    return getAccount(token).then((account) => {
        return fetch(myRequest(`${API_URL}/accounts/${account.id}/activity/${activityId}`, 'GET', token))
            .then((response) => {
                if (response.ok) {
                    return response.json();
                }
                return rejectPromise(response);
            })
            .catch((err) => {
                console.log(err);
                return rejectPromise(err);
            });
    });
};

export const getUserCards = (
    token: string
): Promise<Card[]> => {
    return getAccount(token).then((account) => {
        return fetch(myRequest(`${API_URL}/accounts/${account.id}/cards`, 'GET', token))
            .then((response) => {
                if (response.ok) {
                    return response.json();
                }
                return rejectPromise(response);
            })
            .catch((err) => {
                console.log(err);
                return rejectPromise(err);
            });
    });
};

export const deleteUserCard = (
    userId: string,
    cardId: string,
    token: string
): Promise<Response> => {
    return fetch(
        myRequest(`${API_URL}/accounts/${userId}/cards/${cardId}`, 'DELETE', token)
    )
        .then((response) => {
            if (response.ok) {
                return response.json();
            }
            return rejectPromise(response);
        })
        .catch((err) => {
            console.log(err);
            return rejectPromise(err);
        });
};

export const createUserCard = (card: any, token: string): Promise<Response> => {
    return getAccount(token).then((account) => {
        return fetch(
            myRequest(`${API_URL}/accounts/${account.id}/cards`, 'POST', token),
            {
                body: JSON.stringify(card),
            }
        )
            .then((response) =>
                response.ok ? response.json() : rejectPromise(response)
            )
            .catch((err) => {
                console.log(err);
                return rejectPromise(err);
            });
    });
};

export const createDepositActivity = (
    cardId: number,
    amount: number,
    token: string
) => {
    const activity = {
        amount,
        card_id: cardId
    };

    return getAccount(token).then((account) => {
        return fetch(
            myRequest(`${API_URL}/accounts/${account.id}/deposit`, 'POST', token),
            {
                body: JSON.stringify(activity),
            }
        )
            .then((response) =>
                response.ok ? response.json() : rejectPromise(response)
            )
            .catch((err) => {
                console.log(err);
                return rejectPromise(err);
            });
    });
};

export const createTransferActivity = (
    userId: string,
    token: string,
    origin: string,
    destination: string,
    amount: number,
    description?: string
) => {
    return getAccount(token).then((account) => {
        return fetch(
            myRequest(`${API_URL}/accounts/${account.id}/transactions`, 'POST', token),
            {
                body: JSON.stringify({
                    amount: amount,
                    origin_cvu: origin,
                    destination_cvu: destination,
                    description: description,
                }),
            }
        )
            .then((response) =>
                response.ok ? response.json() : rejectPromise(response)
            )
            .then((response) => {
                return response;
            })
            .catch((err) => {
                console.log(err);
                return rejectPromise(err);
            });
    });
};
