/* eslint-disable jsx-a11y/click-events-have-key-events */
import React, { useEffect, useState } from 'react';

export enum ToolTipEventType {
    click = 'click',
    hover = 'hover',
}

export enum TooltipPosition {
    top = 'top',
    bottom = 'bottom',
}

export interface TooltipProps {
    children?: React.ReactNode;
    message?: string;
    event?: ToolTipEventType;
    position?: TooltipPosition;
    className?: string;
}

const toolTipPositionStyle = {
    top: 'tw-flex-col-reverse tw-bottom-full tw-mb-2',
    bottom: 'tw-flex-col tw-top-full tw-mt-2',
};

export const Tooltip = ({
                            children,
                            message,
                            event = ToolTipEventType.click,
                            position = TooltipPosition.bottom,
                            className = '',
                        }: TooltipProps) => {
    const [isActive, setIsActive] = useState(false);
    const updateActive = () => setIsActive(true);

    useEffect(() => {
        const cancelActive = () => setIsActive(false);
        if (isActive) {
            setTimeout(() => window.addEventListener('click', cancelActive));
            window.addEventListener('scroll', cancelActive);
        }

        return () => {
            removeEventListener('click', cancelActive);
            removeEventListener('scroll', cancelActive);
        };
    }, [isActive]);

    return (
        <div
            className={`tw-relative tw-inline-flex tw-justify-center ${className}`}
            onClick={event === ToolTipEventType.click ? updateActive : undefined}
            onMouseEnter={event === ToolTipEventType.hover ? updateActive : undefined}
            onMouseLeave={
                event === ToolTipEventType.hover ? () => setIsActive(false) : undefined
            }
        >
            <>{children}</>
            {isActive && (
                <div
                    className={`tw-absolute tw-flex tw-items-center tw-transition tw-duration-150 tw-ease-out tw-z-50 ${toolTipPositionStyle[position]}`}
                >
                    <svg
                        className={`tw-relative tw-text-primary ${
                            position === TooltipPosition.bottom
                                ? 'tw-top-px'
                                : 'tw--top-px tw-rotate-180'
                        }`}
                        fill="none"
                        height="8"
                        viewBox="0 0 22 8"
                        width="22"
                        xmlns="http://www.w3.org/2000/svg"
                    >
                        <path
                            d="M11 0L0.607697 7.52632L21.3923 7.52632L11 0Z"
                            fill="currentColor"
                        />
                    </svg>

                    <span
                        className="tw-bg-primary tw-px-4 tw-py-2 tw-text-white tw-rounded tw-inline-block tw-text-center">
            {message}
          </span>
                </div>
            )}
        </div>
    );
};
