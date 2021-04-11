import React from 'react';
import TableOfUsers from "../TableOfUsers/TableOfUsers";
import UserDetails from "../UserDefails/UserDetails";

export const UserForm = ({match}) => (
    <div className="login-wrapper">
        <div className="my-divTable">
            <div className="my-divTableBody">
                <div className="my-divTableRow">
                    <div className="my-divTableCellLeft">&nbsp;</div>
                    <div className="my-divTableCell">
                        <h1>User form</h1>
                    </div>
                    <div className="my-divTableCellRight">&nbsp;</div>
                </div>
            </div>
            <UserDetails id={match.params.id}/>
        </div>
    </div>
);

export default UserForm;