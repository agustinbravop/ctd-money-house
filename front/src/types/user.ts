export interface User {
    firstName: string;
    lastName: string;
    email: string;
    password: string;
    // API: nosotros tenemos 'telephone' en lugar de 'phone'.
    telephone?: string;
    // API: nuestro dni es un string en lugar de un number.
    dni?: string;
    id?: string;
}

export interface UserAccount {
    // API: nosotros tenemos al 'balance' como 'amount'.
    amount: number;
    cvu: string;
    alias: string;
    userId: string;
    id: string;
    name: string;
}
