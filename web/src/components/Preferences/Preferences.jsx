import React from 'react';
import {Link} from "react-router-dom";

export default function Preferences() {
    return (
        <div className="login-wrapper">
            <Link to='/dashboard' key='dashboard'><h2>Dashboard</h2></Link><h2>Preferences</h2>
            <h1>Preferences</h1>
        </div>
    );
}
