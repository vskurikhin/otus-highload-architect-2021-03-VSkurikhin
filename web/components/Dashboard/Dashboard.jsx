import React from 'react';
import {Link} from "react-router-dom";
import TableOfUsers from "../TableOfUsers/TableOfUsers";

export default function Dashboard() {
    return (
        <div className="login-wrapper">
            <h2><Link to='/login' key='login'>Login</Link></h2>
            <h2><Link to='/signin' key='signin'>Sign-in</Link></h2>
            <h2><span style={{color: "#555555"}}>Dashboard</span></h2>
            <div className="my-divTable">
                <div className="my-divTableBody">
                    <div className="my-divTableRow">
                        <div className="my-divTableCellLeft">&nbsp;</div>
                        <div className="my-divTableCell">
                            <TableOfUsers />
                        </div>
                        <div className="my-divTableCellRight">&nbsp;</div>
                    </div>
                </div>
            </div>
        </div>
    );
}
