import * as React from 'react';

const SvgTransferIn = (props: React.SVGProps<SVGSVGElement>): JSX.Element => (
    <svg
        width={50}
        height={50}
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
        {...props}
    >
        <circle cx={25} cy={25} r={24.5} stroke="currentColor" />
        <path
            d="M25.088 27.456c1.597 0 2.838-1.46 2.838-3.191 0-1.732-1.241-3.192-2.838-3.192-1.596 0-2.838 1.46-2.838 3.192 0 1.731 1.242 3.19 2.838 3.19Z"
            stroke="currentColor"
            strokeWidth={0.5}
        />
        <ellipse cx={15.706} cy={23.897} rx={0.971} ry={1.103} fill="currentColor" />
        <ellipse cx={33.824} cy={23.897} rx={0.971} ry={1.103} fill="currentColor" />
        <path
            fillRule="evenodd"
            clipRule="evenodd"
            d="M31.611 15.441H9.56V33.088h31.059V19.581a5.19 5.19 0 0 1-.8.73v11.977H10.359V16.241h21.2v-.065c0-.25.018-.495.052-.735Z"
            fill="currentColor"
        />
        <path
            d="M33.76 16.177h6.47m-6.47 0 2.83-2.942m-2.83 2.942 2.83 2.94"
            stroke="currentColor"
            strokeLinecap="round"
            strokeLinejoin="round"
        />
    </svg>
);

export default SvgTransferIn;