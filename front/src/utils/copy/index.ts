export const copyToClipboard = async (textToCopy: string) => {
    if ('clipboard' in navigator) {
        await navigator.clipboard.writeText(textToCopy);
    } else {
        document.execCommand('copy', true, textToCopy);
    }
};
