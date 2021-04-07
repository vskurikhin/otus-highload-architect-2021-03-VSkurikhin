import './App.css';

import Dashboard from './components/Dashboard/Dashboard';
import Login from "./components/Login/Login";
import Preferences from './components/Preferences/Preferences';
import Signin from "./components/Signin/Signin";

import React, {useState} from 'react';
import {BrowserRouter, Redirect, Route, Switch} from 'react-router-dom';

function main(setToken) {
    return (
        <div className="wrapper">
            <BrowserRouter>
                <Switch>
                    <Route path="/dashboard">
                        <Dashboard/>
                    </Route>
                    <Route path="/preferences">
                        <Preferences/>
                    </Route>
                    <Route path="/signin">
                        <Signin setToken={setToken}/>
                    </Route>
                    <Route path="/login">
                        <Login setToken={setToken}/>
                    </Route>
                    <Redirect from='/' to='/dashboard'/>
                </Switch>
            </BrowserRouter>
        </div>
    );
}

function login(setToken) {
    return (
        <div className="wrapper">
            <BrowserRouter>
                <Switch>
                    <Route path="/signin">
                        <Signin setToken={setToken}/>
                    </Route>
                    <Route path="/login">
                        <Login setToken={setToken}/>
                    </Route>
                    <Redirect from='/' to='/login'/>
                </Switch>
            </BrowserRouter>
        </div>
    );
}

function App() {

    const [token, setToken] = useState();

    if (!token) {
        return login(setToken)
    }
    return main(setToken)
}

export default App;
