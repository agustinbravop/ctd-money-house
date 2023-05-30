import * as React from 'react';

const SvgAdd = (props: React.SVGProps<SVGSVGElement>): JSX.Element => (
    <svg
        width={34}
        height={34}
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
        {...props}
    >
        <circle cx={17} cy={17} r={16.35} stroke="currentColor" strokeWidth={1.3} />
        <path
            d="M16.75 10v14.5M24 17H9.5"
            stroke="currentColor"
            strokeWidth={1.3}
            strokeLinejoin="round"
        />
    </svg>
);

export default SvgAdd;