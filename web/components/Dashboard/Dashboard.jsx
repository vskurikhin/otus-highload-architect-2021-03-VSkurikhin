import React from 'react';
import {Link} from "react-router-dom";

export default function Dashboard() {
    return (
        <div className="login-wrapper">
            <Link to='/login' key='login'><h2>Login</h2></Link><Link to='/signin' key='signin'><h2>Sign-in</h2></Link><h2>Dashboard</h2>
            <h1>Dashboard</h1>
        </div>
    );
}
