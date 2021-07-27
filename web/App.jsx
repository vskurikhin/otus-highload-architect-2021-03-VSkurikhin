import './App.css'

import AppMenu from "./components/AppMenu/AppMenu"
import Error from "./components/Error/Error"
import Login from "./components/Login/Login"
import MyProfile from "./components/MyProfile/MyProfile"
import Signin from "./components/Signin/Signin"
import UserMessage from "./components/UserMessage/UserMessage";

import React from 'react'
import {BrowserRouter, Redirect, Route, Switch} from 'react-router-dom'

const App = () => {
    return (
        <div className="wrapper">
            <BrowserRouter>
                <AppMenu
                    onItemClick={item => this.onItemClick(item)}
                    items={[
                        ["Login", "/login"],
                        ["Sign-in", "/signin"],
                        ["Messages", "/messages", true],
                        ["Profile", "/myprofile", true]
                    ]}
                />
                <Switch>
                    <Route path="/login">
                        <Login/>
                    </Route>
                    <Route path="/signin">
                        <Signin/>
                    </Route>
                    <Route path="/messages">
                        <UserMessage/>
                    </Route>
                    <Route path="/myprofile" component={MyProfile}>
                        <MyProfile/>
                    </Route>
                    <Route path="/error/:msg" component={Error}/>
                    <Redirect from='/' to='/login'/>
                </Switch>
            </BrowserRouter>
        </div>
    )
}
export default App
