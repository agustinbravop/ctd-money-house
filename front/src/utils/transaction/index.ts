import { ActivityType, TransactionType } from '../../types';

export const calculateTransacionType = (
    amount: number,
    type: string
): ActivityType => {
    const isNegative = amount < 0;
    if (type === TransactionType.Transfer) {
        return isNegative ? ActivityType.TRANSFER_OUT : ActivityType.TRANSFER_IN;
    }
    return ActivityType.DEPOSIT;
};
