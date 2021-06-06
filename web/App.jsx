import './App.css'

import AddNews from "./components/AddNews/AddNews";
import AppMenu from "./components/AppMenu/AppMenu"
import Error from "./components/Error/Error"
import Login from "./components/Login/Login"
import MyProfile from "./components/MyProfile/MyProfile"
import NewsList from './components/NewsList/NewsList'
import Signin from "./components/Signin/Signin"
import UserForm from './components/UserForm/UserForm'
import UserList from './components/UserList/UserList'
import UserSearch from "./components/UserSearch/UserSearch";

import React, {useEffect, useRef, useState} from 'react'
import {BrowserRouter, Redirect, Route, Switch} from 'react-router-dom'
import ChatWindow from "./components/ChatWindow/ChatWindow";
import ChatEntry from "./components/ChatEntry/ChatEntry";

const App = () => {
    return (
        <div className="wrapper">
            <BrowserRouter>
                <AppMenu
                    onItemClick={item => this.onItemClick(item)}
                    items={[
                        ["Login", "/login"],
                        ["Sign-in", "/signin"],
                        ["Add news", "/addnews", true, true],
                        ["News list", "/newslist", true],
                        // ["User list", "/userlist", true],
                        ["User search", "/usersearch", true],
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
                    <Route path="/addnews">
                        <AddNews/>
                    </Route>
                    <Route path="/newslist">
                        <NewsList/>
                    </Route>
                    {/*<Route path="/userlist">*/}
                    {/*    <UserList/>*/}
                    {/*</Route>*/}
                    <Route path="/usersearch">
                        <UserSearch/>
                    </Route>
                    <Route path="/userform/:id" component={UserForm}/>
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
