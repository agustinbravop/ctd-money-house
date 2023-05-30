const visaRegex = /^4\d{3}-?\d{4}-?\d{4}-?\d{4}$/;
const mastercardRegex = /^5[1-5]\d{2}-?\d{4}-?\d{4}-?\d{4}$/;

export const isVisa = (number: string) => visaRegex.test(number);
export const isMastercard = (number: string) => mastercardRegex.test(number);