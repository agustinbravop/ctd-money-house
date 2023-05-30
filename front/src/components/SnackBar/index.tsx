import React, { useState } from 'react';
import Snackbar from '@mui/material/Snackbar';

export interface ISnackBar {
    duration: number;
    message: string;
    type?: 'primary' | 'success' | 'error';
}

const bgType = {
    primary: 'tw-bg-primary',
    success: 'tw-bg-success',
    error: 'tw-bg-error',
};

export const SnackBar = ({
                             duration,
                             message,
                             type = 'primary',
                         }: ISnackBar) => {
    const [open, setOpen] = useState<boolean>(true);

    const closeNotification = () => {
        setOpen(false);
    };

    return (
        <Snackbar
            open={open}
            autoHideDuration={duration}
            onClose={closeNotification}
        >
            <div className={`${bgType[type]} tw-p-4 tw-rounded`}>{message}</div>
        </Snackbar>
    );
};
