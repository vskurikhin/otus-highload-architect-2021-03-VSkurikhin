
import './App.css'

import AppMenu from "./components/Menu/AppMenu"
import Error from "./components/Error/Error";
import Login from "./components/Login/Login"
import Signin from "./components/Signin/Signin"
import UserForm from './components/UserForm/UserForm'
import UserList from './components/UserList/UserList'

import React, {useState} from 'react'
import {BrowserRouter, Redirect, Route, Switch} from 'react-router-dom'

const main = setToken => (
    <div className="wrapper">
        <BrowserRouter>
            <AppMenu disabledUserList={false}/>
            <Switch>
                <Route path="/userlist">
                    <UserList/>
                </Route>
                <Route path="/userform/:id" component={UserForm} />
                <Route path="/signin">
                    <Signin setToken={setToken}/>
                </Route>
                <Route path="/login">
                    <Login setToken={setToken}/>
                </Route>
                <Route path="/error/:msg" component={Error} />
                <Redirect from='/' to='/userlist'/>
            </Switch>
        </BrowserRouter>
    </div>
)

const login = setToken => (
    <div className="wrapper">
        <BrowserRouter>
            <AppMenu disabledUserList={true}/>
            <Switch>
                <Route path="/signin">
                    <Signin setToken={setToken}/>
                </Route>
                <Route path="/login">
                    <Login setToken={setToken}/>
                </Route>
                <Route path="/error/:msg" component={Error} />
                <Redirect from='/' to='/login'/>
            </Switch>
        </BrowserRouter>
    </div>
)

function App() {

    const [token, setToken] = useState();

    if (token) {
        return main(setToken)
    }
    return login(setToken)

}

export default App
