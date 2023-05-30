import * as React from 'react';

const SvgCopy = (props: React.SVGProps<SVGSVGElement>): JSX.Element => (
    <svg
        width={32}
        height={32}
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
        {...props}
    >
        <path
            d="M28 10v18H10V10h18Zm0-2H10a2 2 0 0 0-2 2v18a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V10a2 2 0 0 0-2-2Z"
            fill="currentColor"
        />
        <path d="M4 18H2V4a2 2 0 0 1 2-2h14v2H4v14Z" fill="currentColor" />
    </svg>
);

export default SvgCopy;
