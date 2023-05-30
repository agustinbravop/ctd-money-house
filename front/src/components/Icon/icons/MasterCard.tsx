import * as React from 'react';

const SvgMastercard = (props: React.SVGProps<SVGSVGElement>): JSX.Element => (
    <svg
        width={34}
        height={34}
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
        {...props}
    >
        <path d="M20.238 11.588h-6.495v11.668h6.495V11.588Z" fill="#FF5A00" />
        <path
            d="M14.175 17.422A7.443 7.443 0 0 1 17 11.588a7.418 7.418 0 0 0-12 5.834 7.418 7.418 0 0 0 12 5.834 7.407 7.407 0 0 1-2.825-5.834Z"
            fill="#EB001B"
        />
        <path
            d="M29 17.422a7.418 7.418 0 0 1-12 5.834 7.382 7.382 0 0 0 2.825-5.834A7.443 7.443 0 0 0 17 11.588 7.37 7.37 0 0 1 21.576 10C25.68 10 29 13.341 29 17.422Z"
            fill="#F79E1B"
        />
        <circle cx={17} cy={17} r={16.5} stroke="currentColor" />
    </svg>
);

export default SvgMastercard;
