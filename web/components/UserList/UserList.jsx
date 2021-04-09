import React from 'react';
import TableOfUsers from "../TableOfUsers/TableOfUsers";
import {useHistory} from "react-router-dom";

export const UserList = props => {

    const history = useHistory();

    try {
        return (
            <div className="login-wrapper">
                <div className="my-divTable">
                    <div className="my-divTableBody">
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <h1>User list</h1>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <TableOfUsers {...props}/>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                    </div>
                </div>
            </div>
        );
    } catch (e) {
        console.debug(e);
        history.push('/login');
        return <div/>;
    }
}

export default UserList;
