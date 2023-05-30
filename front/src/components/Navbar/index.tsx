import React, { useEffect } from 'react';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Button from '@mui/material/Button';
import Avatar from '@mui/material/Avatar';
import { Icon } from '../Icon';
import { Link, useLocation } from 'react-router-dom';
import { ROUTES } from '../../constants';
import { useUserInfo } from '../../hooks';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
import Fade from '@mui/material/Fade';
import { useAuth } from '../../hooks/useAuth';
import { Skeleton, SkeletonVariant } from '../Skeleton';

const stringAvatar = (name: string) => {
    if (name.length === 0) return { children: name };
    return {
        children: `${name.split(' ')[0][0]}${name.split(' ')[1][0]}`,
    };
};

export const Navbar = ({ isAuthenticated = false }) => {
    const { user, loading } = useUserInfo();
    const [fullName, setFullName] = React.useState('');
    const location = useLocation();
    const isLogin = location.pathname === ROUTES.LOGIN && !isAuthenticated;
    const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
    const { logout } = useAuth();

    const handleClick = (event: React.MouseEvent<HTMLElement>) => {
        setAnchorEl(event.currentTarget);
    };

    const handleClose = () => {
        setAnchorEl(null);
    };

    const open = Boolean(anchorEl);

    useEffect(() => {
        if (user) {
            setFullName(`${user.firstName} ${user.lastName}`);
        }
    }, [user]);

    return (
        <div className="tw-w-full tw-fixed tw-z-50 print:tw-hidden">
            <AppBar
                className="tw-px-10 !tw-text-neutral-gray-100 tw-border-b tw-border-neutral-blue-100 tw-backdrop-blur tw-shadow-none"
                style={{ background: 'transparent' }}
            >
                <Toolbar className="tw-flex tw-px-0 tw-justify-between">
                    <Link to={ROUTES.HOME}>
                        <Icon className="tw-text-primary" type="digital-house" />
                    </Link>
                    {!isAuthenticated ? (
                        <Link to={isLogin ? ROUTES.REGISTER : ROUTES.LOGIN}>
                            <Button variant="contained">
                                {isLogin ? 'Crear cuenta' : 'Iniciar Sesión'}
                            </Button>
                        </Link>
                    ) : (
                        <div>
                            {loading ? (
                                <Skeleton variant={SkeletonVariant.SQUARE} />
                            ) : (
                                <>
                                    <Button
                                        className="tw-flex tw-items-center tw-gap-x-2"
                                        id="fade-button"
                                        aria-controls={open ? 'fade-menu' : undefined}
                                        aria-haspopup="true"
                                        aria-expanded={open ? 'true' : undefined}
                                        onClick={handleClick}
                                    >
                                        <Avatar
                                            className="tw-bg-primary tw-rounded-xl"
                                            {...stringAvatar(fullName)}
                                        />
                                        Hola, {fullName}
                                    </Button>
                                    <Menu
                                        id="fade-menu"
                                        MenuListProps={{
                                            'aria-labelledby': 'fade-button',
                                        }}
                                        anchorEl={anchorEl}
                                        open={open}
                                        onClose={handleClose}
                                        TransitionComponent={Fade}
                                    >
                                        <MenuItem onClick={logout}>Cerrar Sesión</MenuItem>
                                    </Menu>
                                </>
                            )}
                        </div>
                    )}
                </Toolbar>
            </AppBar>
        </div>
    );
};
