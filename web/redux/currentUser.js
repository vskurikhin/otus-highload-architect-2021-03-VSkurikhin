/*
 * This file was last modified at 2021.03.21 17:13 by Victor N. Skurikhin.
 * This is free and unencumbered software released into the public domain.
 * For more information, please refer to <http://unlicense.org>
 * currentUser.js
 * $Id$
 */

import {GET} from "../lib/consts"
import throwResultCode from "../lib/throwResultCode"
import deleteCookie from "../lib/deleteCookie";

export const getProfileFetch = () => {
    return dispatch => {

        const getResult = result => {

            if (result.Code && result.Message) {
                window.sessionStorage.removeItem("username")
                deleteCookie("acs_jwt")
                throwResultCode(result)
            }
            const {Id, Username} = result
            if (Username) {
                window.sessionStorage.setItem("username", Username)
                dispatch(loginUser({Id, Username}))
            }
        }
        const getError = error => {
            throw error
        }
        return fetch("/profile", GET)
            .then(res => res.json())
            .then(getResult, getError)
    }
}

export const loginUser = ({Id, Username}) => ({
    type: 'LOGIN_USER',
    payload: {Id, Username}
})

export const logoutUser = () => ({
    type: 'LOGOUT_USER'
})
