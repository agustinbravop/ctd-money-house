const customTheme = {
    screens: {
        sm: '320px',
        md: '720px',
        lg: '1024px',
        xl: '1280px',
        '2xl': '1920px',
    },
    extend: {
        colors: {
            primary: {
                DEFAULT: 'var(--primary)',
            },
            secondary: {
                DEFAULT: 'var(--secondary)',
            },
            background: {
                DEFAULT: 'var(--background)',
            },
            error: {
                DEFAULT: 'var(--error)',
            },
            success: {
                DEFAULT: 'var(--success)',
            },
            white: {
                DEFAULT: 'var(--white)',
            },
            neutral: {
                'gray-100': 'var(--neutral-gray-100)',
                'gray-300': 'var(--neutral-gray-300)',
                'gray-500': 'var(--neutral-gray-500)',
                'blue-100': 'var(--neutral-blue-100)',
            },
        },
        fontSize: {
            '8xl': 'var(--8xl)',
            '7xl': 'var(--7xl)',
            '6xl': 'var(--6xl)',
            '4xl': 'var(--4xl)',
            h2xl: 'var(--h2xl)',
            '3xl': 'var(--3xl)',
            subtext: 'var(--subtext)',
            '2xl': 'var(--2xl)',
            xl: 'var(--xl)',
            lg: 'var(--lg)',
            base: 'var(--base)',
            sm: 'var(--sm)',
            xs: 'var(--xs)',
        },
        borderRadius: {
            image: '1.25rem',
        },
        spacing: {
            icon: '0.375rem',
            4.5: '1.125rem',
            18: '4.5rem',
            119: '29.75rem',
            150: '37.5rem',
            152: '38rem',
            192: '48rem',
        },
        gridTemplateColumns: {
            'auto-fill': 'repeat(auto-fill, minmax(412px, 1fr))',
        },
        opacity: {
            38: '.38',
        },
        maxHeight: {
            '4/5': '80%',
        },
        width: {
            156: '39rem',
        },
        maxWidth: {
            'container-md': '748px',
            'container-lg': '1268px',
        },
        aspectRatio: {
            '5/4': '5 / 4',
        },
        lineHeight: {
            CTA: 1.125,
            heading: '72px',
        },
    },
};

exports.theme = customTheme;
