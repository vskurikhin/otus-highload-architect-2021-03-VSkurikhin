import './Signin.css';

import React, {useState} from 'react';
import PropTypes from 'prop-types';
import {Link} from "react-router-dom";

async function signinUser(credentials) {
    return fetch('http://localhost:6060/signin', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(credentials)
    })
        .then(data => data.json())
}

export default function Signin({setToken}) {

    const [username, setUserName] = useState();
    const [password, setPassword] = useState();

    const handleSubmit = async e => {
        e.preventDefault();
        const token = await signinUser({
            username,
            password
        });
        setToken(token);
    }

    return (
        <div className="signin-wrapper">
            <Link to='/login' key='login'><h2>Login</h2></Link><h2>Sign-in</h2>
            <h1>For register Sign in please</h1>
            <form onSubmit={handleSubmit}>
                <label>
                    <p>Username</p>
                    <input type="text" onChange={e => setUserName(e.target.value)}/>
                </label>
                <label>
                    <p>Password</p>
                    <input type="password" name="password" onChange={e => setPassword(e.target.value)}/>
                </label>
                <div>
                    <button type="submit">Submit</button>
                </div>
            </form>
        </div>
    )
}

Signin.propTypes = {
    setToken: PropTypes.func.isRequired
}