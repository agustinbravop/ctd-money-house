import React from 'react';
import { Navbar } from '../Navbar';
import { Footer } from '../Footer';
import { Sidebar } from '../Sidebar';

interface NavbarProps {
    children: React.ReactNode;
    isAuthenticated?: boolean;
}

const navbarHeight = '4rem';

export const Layout = ({
                           children,
                           isAuthenticated = false,
                       }: NavbarProps): JSX.Element => {
    return (
        <>
            <div
                style={{
                    marginTop: navbarHeight,
                }}
                className="tw-flex print:!tw-mt-0"
            >
                <Navbar isAuthenticated={isAuthenticated} />
                {isAuthenticated && <Sidebar />}

                <main
                    className="tw-flex tw-flex-col tw-flex-1  tw-flex-wrap tw-overflow-auto print:!tw-p-0"
                    style={
                        !isAuthenticated
                            ? {
                                minHeight: 'calc(100vh - 8rem)',
                            }
                            : {
                                paddingBottom: '3rem',
                            }
                    }
                >
                    {children}
                </main>
            </div>
            <Footer />
        </>
    );
};
