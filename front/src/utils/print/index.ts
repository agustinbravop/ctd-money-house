export const printPage = (): void =>
    typeof window === 'undefined' ? undefined : window.print();
