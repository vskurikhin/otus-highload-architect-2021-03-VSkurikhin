
import {combineReducers} from 'redux'
import {routerReducer} from 'react-router-redux'

import currentUser from "./currentUser"

const rootReducer = combineReducers({
    routing: routerReducer,
    currentUser: currentUser,
})

export default rootReducer