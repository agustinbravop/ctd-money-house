import { useContext } from 'react';
import { userInfoContext } from './../context/User/';

export const useUserInfo = () => {
    return useContext(userInfoContext);
};
