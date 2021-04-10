import React from 'react';
import TableOfUsers from "../TableOfUsers/TableOfUsers";
import UserDefails from "../UserDefails/UserDefails";

export const UserForm = ({ match } )=> {

    console.log(match);

    return (
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
                <UserDefails id={match.params.id}/>
            </div>
        </div>
    );
}

export default UserForm;