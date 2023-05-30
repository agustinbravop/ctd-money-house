const CURRENCY_FORMATTER = (locales: string, currency: string) =>
    new Intl.NumberFormat(locales, {
        style: 'currency',
        currency: `${currency}`,
        minimumFractionDigits: 0,
    });

export const formatCurrency = (
    locales: string,
    currency: string,
    amount: number
) => {
    return CURRENCY_FORMATTER(locales, currency).format(amount);
};
