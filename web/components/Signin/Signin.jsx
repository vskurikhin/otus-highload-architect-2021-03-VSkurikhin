import './Signin.css';

import React, {useState} from 'react';
import PropTypes from 'prop-types';
import {Link, useHistory} from "react-router-dom";

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

    const history = useHistory();
    const [username, setUserName] = useState();
    const [password, setPassword] = useState();
    const [name, setName] = useState();
    const [surname, setSurname] = useState();
    const [age, setAge] = useState();
    const [sex, setSex] = useState();
    const [city, setCity] = useState();
    const [interests, setInterests] = useState();

    const handleSubmit = async e => {
        e.preventDefault();
        const token = await signinUser({
            username,
            password,
            name,
            surname,
            age,
            sex,
            city,
            interests
        });
        setToken(token);
        if (token){
            history.push('/dashboard');
        }
    }

    return (
        <div className="signin-wrapper">
            <Link to='/login' key='login'><h2>Login</h2></Link><h2>Sign-in</h2><h2>Dashboard</h2>
            <form onSubmit={handleSubmit}>
                <div className="my-divTable">
                    <div className="my-divTableBody">
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <h1>For register Sign in please</h1>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p>Username</p>
                                <input type="text" onChange={e => setUserName(e.target.value)}/>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p>Password</p>
                                <input type="password" name="password" onChange={e => setPassword(e.target.value)}/>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p>Firstname</p>
                                <input type="text" name="name" onChange={e => setName(e.target.value)}/>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p>Surname</p>
                                <input type="text" name="surname" onChange={e => setSurname(e.target.value)}/>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p>Age</p>
                                <input type="text" name="age" onChange={e => setAge(e.target.value)}/>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p>Sex</p>
                                <input type="text" name="sex" onChange={e => setSex(e.target.value)}/>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p>City</p>
                                <input type="text" name="city" onChange={e => setCity(e.target.value)}/>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p>Interests</p>
                                <textarea rows="10" cols="42" name="interests" onChange={e => setInterests(e.target.value)}/>
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

Signin.propTypes = {
    setToken: PropTypes.func.isRequired
}