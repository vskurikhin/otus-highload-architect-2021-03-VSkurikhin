import React from 'react';
import {Link} from "react-router-dom";

export default function Dashboard() {
    return (
        <div className="login-wrapper">
            <h2>Dashboard</h2><Link to='/preferences' key='preferences'><h2>Preferences</h2></Link>
            <h1>Dashboard</h1>
        </div>
    );
}
