
import './Login.css'

import PropTypes from 'prop-types'
import React, {useState} from 'react'
import {Input} from "semantic-ui-react"
import {useHistory} from "react-router-dom"

import deleteCookie from "../../lib/deleteCookie";
import {POST} from "../../lib/consts";

async function loginUser(credentials) {
    return fetch('/login', {
        body: JSON.stringify(credentials),
        ...POST
    }).then(data => data.json())
}

export default function Login({setToken}) {

    const history = useHistory()
    const [username, setUserName] = useState()
    const [password, setPassword] = useState()

    const handleSubmit = async e => {
        e.preventDefault();
        deleteCookie("acs_jwt")
        const token = await loginUser({
            username,
            password
        })
        setToken(token)
        if (token){
            history.push('/userlist')
        }
    }

    return (
        <div className="login-wrapper">
            <form onSubmit={handleSubmit}>
                <div className="my-divTable">
                    <div className="my-divTableBody">
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <h1>Please Log In</h1>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p className="my-p-label">Username:</p>
                                <Input placeholder='Username...' onChange={e => setUserName(e.target.value)}/>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p className="my-p-label">Password:</p>
                                <Input
                                    type="password"
                                    name="password"
                                    placeholder='password...'
                                    onChange={e => setPassword(e.target.value)}
                                />
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <button type="submit">Submit</button>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                    </div>
                </div>
            </form>
        </div>
    )
}

Login.propTypes = {
    setToken: PropTypes.func.isRequired
}
