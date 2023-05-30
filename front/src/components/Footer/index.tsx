import React from 'react';

export const Footer = () => {
    const currentYear = new Date().getFullYear();
    return (
        <footer
            className="tw-flex tw-items-center tw-h-16 tw-px-10 tw-w-full tw-bottom-0 tw-border-t tw-border-neutral-blue-100 tw-text-sm print:tw-hidden">
            <p>{currentYear && `Copyright © ${currentYear}`} Digital Money House</p>
        </footer>
    );
};
