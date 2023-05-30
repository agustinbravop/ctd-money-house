import React from 'react';
import SvgDigitalHouse from './icons/DigitalHouse';
import SvgCheck from './icons/Check';
import SvgArrowRight from './icons/ArrowRight';
import SvgAdd from './icons/Add';
import SvgCopy from './icons/Copy';
import SvgCreditCard from './icons/CreditCard';
import SvgIncome from './icons/Income';
import SvgMastercard from './icons/MasterCard';
import SvgVisa from './icons/Visa';
import SvgTransferIn from './icons/TransferIn';
import SvgTransferOut from './icons/TransferOut';
import SvgUser from './icons/User';
import SvgWithdraw from './icons/Withdraw';
import SvgEdit from './icons/Edit';

export type IconType =
    | 'digital-house'
    | 'check'
    | 'arrow-right'
    | 'add'
    | 'copy'
    | 'credit-card'
    | 'deposit'
    | 'mastercard'
    | 'visa'
    | 'transfer-in'
    | 'transfer-out'
    | 'user'
    | 'withdraw'
    | 'edit';

export interface IconProps extends React.SVGProps<SVGSVGElement> {
    type: IconType;
}

const icons: Record<IconType,
    (props: React.SVGProps<SVGSVGElement>) => JSX.Element> = {
    'digital-house': (props) => <SvgDigitalHouse role="img" {...props} />,
    check: (props) => <SvgCheck role="img" {...props} />,
    'arrow-right': (props) => <SvgArrowRight role="img" {...props} />,
    add: (props) => <SvgAdd role="img" {...props} />,
    copy: (props) => <SvgCopy role="img" {...props} />,
    'credit-card': (props) => <SvgCreditCard role="img" {...props} />,
    deposit: (props) => <SvgIncome role="img" {...props} />,
    mastercard: (props) => <SvgMastercard role="img" {...props} />,
    visa: (props) => <SvgVisa role="img" {...props} />,
    'transfer-in': (props) => <SvgTransferIn role="img" {...props} />,
    'transfer-out': (props) => <SvgTransferOut role="img" {...props} />,
    user: (props) => <SvgUser role="img" {...props} />,
    withdraw: (props) => <SvgWithdraw role="img" {...props} />,
    edit: (props) => <SvgEdit role="img" {...props} />,
};

export const Icon = ({ type, ...restProps }: IconProps): JSX.Element | null => {
    const icon = icons[type];
    if (icon) {
        return icon(restProps);
    }
    return null;
};
