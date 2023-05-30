import { Transaction } from '../../types';

export const sortByDate = (
    transactions: Transaction[],
    mode = 'desc'
): Transaction[] => {
    if (mode === 'desc' && transactions.length > 0) {
        return transactions.sort((a, b) => {
            return new Date(b.transaction_date).getTime() - new Date(a.transaction_date).getTime();
        });
    }
    if (mode === 'asc' && transactions.length > 0) {
        return transactions.sort((a, b) => {
            return new Date(a.transaction_date).getTime() - new Date(b.transaction_date).getTime();
        });
    }

    return transactions;
};
