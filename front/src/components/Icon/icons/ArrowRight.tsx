import * as React from 'react';

const ArrowRight = (props: React.SVGProps<SVGSVGElement>): JSX.Element => (
    <svg
        width={22}
        height={24}
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
        {...props}
    >
        <path
            d="M2 10.5a1.5 1.5 0 0 0 0 3v-3Zm19.06 2.56a1.5 1.5 0 0 0 0-2.12l-9.545-9.547a1.5 1.5 0 1 0-2.122 2.122L17.88 12l-8.486 8.485a1.5 1.5 0 1 0 2.122 2.122l9.546-9.546ZM2 13.5h18v-3H2v3Z"
            fill="currentColor"
        />
    </svg>
);

export default ArrowRight;
