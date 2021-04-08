import './Login.css';

import PropTypes from 'prop-types';
import React, {useState} from 'react';
import {Input} from "semantic-ui-react";
import {Link, useHistory} from "react-router-dom";

async function loginUser(credentials) {
    return fetch('/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(credentials)
    })
        .then(data => data.json())
}

export default function Login({setToken}) {

    const history = useHistory();
    const [username, setUserName] = useState();
    const [password, setPassword] = useState();

    const handleSubmit = async e => {
        e.preventDefault();
        const token = await loginUser({
            username,
            password
        });
        setToken(token);
        if (token){
            history.push('/dashboard');
        }
    }

    return (
        <div className="login-wrapper">
            <h2><span style={{color: "#555555"}}>Login</span></h2>
            <h2><Link to='/signin' key='signin'>Sign-in</Link></h2>
            <h2>Dashboard</h2>
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