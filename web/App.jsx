import './App.css'

import Error from "./components/Error/Error"
import HeaderMenu from "./components/AppMenu/AppMenu"
import Login from "./components/Login/Login"
import MyProfile from "./components/MyProfile/MyProfile"
import Signin from "./components/Signin/Signin"
import UserForm from './components/UserForm/UserForm'
import UserList from './components/UserList/UserList'

import React from 'react'
import {BrowserRouter, Redirect, Route, Switch} from 'react-router-dom'

const App = () => (
    <div className="wrapper">
        <BrowserRouter>
            <HeaderMenu
                onItemClick={item => this.onItemClick(item)}
                items={[
                    ["Login", "/login"],
                    ["Sign-in", "/signin"],
                    ["User list", "/userlist", true],
                    ["Profile", "/myprofile", true]
                ]}
            />
            <Switch>
                <Route path="/userlist">
                    <UserList/>
                </Route>
                <Route path="/userform/:id" component={UserForm}/>
                <Route path="/signin">
                    <Signin/>
                </Route>
                <Route path="/login">
                    <Login/>
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

export default App
