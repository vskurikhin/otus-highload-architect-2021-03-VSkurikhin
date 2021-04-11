
import React from "react";
import {Input} from "semantic-ui-react";

export const UserInterests = props => {

    const {interests} = props

    let isArray = false
    if (typeof Array.isArray === 'undefined') {
        Array.isArray = function (obj) {
            isArray = Object.prototype.toString.call(obj) === '[object Array]'
        }
    }

    if (Array.isArray(interests) || isArray) {
        if (interests.length < 1) {
            return <div/>
        } else {
            return interests.map((interest) => (
                <div className="my-divTableRow" key={interest}>
                    <div className="my-divTableCellLeft">&nbsp;</div>
                    <div className="my-divTableCell">
                        <Input value={interest} disabled={true}/>
                    </div>
                    <div className="my-divTableCellRight">&nbsp;</div>
                </div>
            ))
        }
    } else {
        return <div/>
    }
}

export default UserInterests
