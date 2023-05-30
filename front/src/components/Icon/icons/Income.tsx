import * as React from 'react';

const SvgIncome = (props: React.SVGProps<SVGSVGElement>): JSX.Element => (
    <svg
        width={51}
        height={51}
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
        {...props}
    >
        <circle cx={25.25} cy={25.25} r={24.75} stroke="currentColor" />
        <path
            d="M25.25 25.734c1.518 0 2.725-1.299 2.725-2.87 0-1.573-1.207-2.872-2.726-2.872-1.518 0-2.725 1.3-2.725 2.871 0 1.572 1.207 2.871 2.725 2.871Z"
            stroke="currentColor"
            strokeWidth={0.5}
        />
        <ellipse cx={16.276} cy={22.536} rx={0.928} ry={0.983} fill="currentColor" />
        <ellipse cx={33.605} cy={22.536} rx={0.928} ry={0.983} fill="currentColor" />
        <path
            fillRule="evenodd"
            clipRule="evenodd"
            d="M11.197 15.8h28.106v14.127h-9.52c.107.256.195.524.262.8h10.058V15H10.397V30.727H20.456c.067-.276.154-.544.261-.8h-9.52V15.8Z"
            fill="currentColor"
        />
        <path
            d="M25.25 29.416v6.553m2.475-3.686-2.476-2.867-2.475 2.867"
            stroke="currentColor"
            strokeLinecap="round"
            strokeLinejoin="round"
        />
    </svg>
);

export default SvgIncome;
