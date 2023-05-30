import * as React from 'react';

const SvgWithdraw = (props: React.SVGProps<SVGSVGElement>): JSX.Element => (
    <svg
        width={50}
        height={50}
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
        {...props}
    >
        <circle cx={25} cy={25} r={24.5} stroke="currentColor" />
        <path
            d="M25 25.337c1.504 0 2.7-1.288 2.7-2.846 0-1.557-1.196-2.845-2.7-2.845-1.506 0-2.702 1.288-2.702 2.845 0 1.558 1.196 2.846 2.701 2.846Z"
            stroke="currentColor"
            strokeWidth={0.5}
        />
        <ellipse cx={16.115} cy={22.167} rx={0.919} ry={0.973} fill="currentColor" />
        <ellipse cx={33.272} cy={22.167} rx={0.919} ry={0.973} fill="currentColor" />
        <path
            fillRule="evenodd"
            clipRule="evenodd"
            d="M11.094 15.506h27.812v13.97h-9.42c.107.257.195.525.262.8h9.958V14.707H10.294V30.276H20.253c.067-.275.155-.543.262-.8h-9.421v-13.97Z"
            fill="currentColor"
        />
        <path
            d="M25 35.467V28.98m-2.452 3.65L25 35.467l2.451-2.838"
            stroke="currentColor"
            strokeLinecap="round"
            strokeLinejoin="round"
        />
    </svg>
);

export default SvgWithdraw;
