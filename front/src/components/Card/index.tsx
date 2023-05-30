import React from 'react';
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';

export interface CardCustomProps {
    content: React.ReactNode;
    actions?: React.ReactNode;
    className?: string;
}

export const CardCustom = ({
                               className,
                               content,
                               actions,
                           }: CardCustomProps) => {
    return (
        <Card
            className={`tw-p-10 tw-max-w-2xl tw-mx-auto tw-mt-12 tw-bg-background tw-text-neutral-gray-100 tw-border-2 tw-border-neutral-blue-100 tw-rounded-lg ${className}`}
            variant="outlined"
        >
            <CardContent className="tw-flex tw-flex-col tw-p-0">
                {content}
            </CardContent>
            {actions && (
                <CardActions className="tw-flex tw-flex-wrap tw-justify-evenly tw-p-0">
                    {actions}
                </CardActions>
            )}
        </Card>
    );
};
