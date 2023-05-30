import React from 'react';
import { default as CustomSkeleton } from '@mui/material/Skeleton';

export enum SkeletonVariant {
    RECORD_LIST = 'record-list',
    SQUARE = 'square',
}

export interface SkeletonProps {
    variant: SkeletonVariant;
    numberOfItems?: number;
    className?: string;
}

export const Skeleton = ({
                             variant = SkeletonVariant.RECORD_LIST,
                             numberOfItems = 10,
                             className = '',
                         }: SkeletonProps) => {
    return (
        <>
            {variant === SkeletonVariant.RECORD_LIST && (
                <div className={`tw-w-full ${className}`}>
                    {[...Array(numberOfItems)].map((_, index) => (
                        <div
                            key={index}
                            className="tw-flex tw-items-center tw-w-full tw-py-5 tw-mb-0.5"
                        >
                            <CustomSkeleton variant="rectangular" height={48} width="100%" />
                        </div>
                    ))}
                </div>
            )}
            {variant === SkeletonVariant.SQUARE && (
                <div className={`tw-w-full ${className}`}>
                    <CustomSkeleton variant="rectangular" height={48} width={128} />
                </div>
            )}
        </>
    );
};
