/**
 * Created by James on 2017-06-25.
 */
import { CREATE_SESSION } from "../actions/session";

const initialState = {

};

const userReducer = (state = initialState, action) => {
    switch (action.type) {
        case CREATE_SESSION:
            // We reset the state to only contain a username and token
            return {
                username: action.username,
                token: action.token
            };
        default:
            return state;
    }
};

export default userReducer;
