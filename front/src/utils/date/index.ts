export const formatDateFromString = (date: string) => {
    const newDate = new Date(date);

    const formatDate =
        newDate.getDate() +
        '-' +
        (newDate.getMonth() + 1) +
        '-' +
        newDate.getFullYear();

    return formatDate;
};
