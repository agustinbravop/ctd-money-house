import { userActionTypes } from './types';
import { UserInfoState } from './';

interface UserAction {
    type: userActionTypes;
    payload: any;
}

export default (state: UserInfoState, action: UserAction) => {
    switch (action.type) {
        case userActionTypes.SET_USER:
            return {
                ...state,
                user: action.payload,
            };
        case userActionTypes.SET_USER_LOADING:
            return {
                ...state,
                loading: action.payload,
            };
        default:
            return state;
    }
};
