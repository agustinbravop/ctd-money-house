import * as React from 'react';

const SvgTransferOut = (props: React.SVGProps<SVGSVGElement>): JSX.Element => (
    <svg
        width={50}
        height={50}
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
        {...props}
    >
        <circle cx={25} cy={25} r={24.5} stroke="currentColor" />
        <path
            d="M25.088 27.456c1.596 0 2.838-1.46 2.838-3.191 0-1.732-1.242-3.192-2.838-3.192-1.597 0-2.838 1.46-2.838 3.192 0 1.731 1.241 3.19 2.838 3.19Z"
            stroke="currentColor"
            strokeWidth={0.5}
        />
        <ellipse cx={15.706} cy={23.897} rx={0.971} ry={1.103} fill="currentColor" />
        <ellipse cx={33.823} cy={23.897} rx={0.971} ry={1.103} fill="currentColor" />
        <path
            fillRule="evenodd"
            clipRule="evenodd"
            d="M31.611 15.441H9.56V33.088h31.058V19.581a5.19 5.19 0 0 1-.8.731v11.976H10.36V16.241h21.2v-.064c0-.25.018-.496.052-.736Z"
            fill="currentColor"
        />
        <path
            d="M40.23 16.177h-6.471m6.47 0-2.83 2.94m2.83-2.94-2.83-2.942"
            stroke="currentColor"
            strokeLinecap="round"
            strokeLinejoin="round"
        />
    </svg>
);

export default SvgTransferOut;
