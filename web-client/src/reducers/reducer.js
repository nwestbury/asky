/**
 * Created by James on 2017-06-18.
 */
import { combineReducers } from 'redux'
import userReducer from './userReducer';

const reducer = combineReducers({
    user: userReducer,
});

export default reducer;
