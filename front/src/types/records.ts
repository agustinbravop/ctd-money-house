export enum TransactionType {
    Transfer = 'egreso',
    Deposit = 'ingreso',
}

export interface Transaction {
    amount: number;
    description?: string;
    transaction_date: string;
    id: string;
    transaction_type: TransactionType;
    origin_cvu?: string;
    destination_cvu?: string;
}

export interface Card {
    card_number: string;
    name: string;
    type: string;
    id: string;
}

export interface Account {
    name: string;
    origin: string;
}

export enum ActivityType {
    TRANSFER_IN = 'egreso',
    TRANSFER_OUT = 'egreso',
    DEPOSIT = 'ingreso',
}
