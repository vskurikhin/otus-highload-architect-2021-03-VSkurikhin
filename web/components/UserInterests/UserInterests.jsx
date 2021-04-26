import PropTypes from 'prop-types'
import React, {useEffect, useState} from "react";
import {Input} from "semantic-ui-react";

export default function UserInterests(props) {

    const [disabled, setDisabled] = useState()
    const [textInterests, setTextInterests] = useState()

    const setInterests2 = values => {
        if (Array.isArray(values)) {
            setTextInterests(values.join('\n'))
        }
    }

    const setTextInterests2 = values => {
        setTextInterests(values)
        props.setInterests(values.split('\n'))
    }

    useEffect(() => setDisabled(props.disabled), [props.disabled])
    useEffect(() => setInterests2(props.interests), [props.interests])

    let isArray = false
    if (typeof Array.isArray === 'undefined') {
        Array.isArray = function (obj) {
            isArray = Object.prototype.toString.call(obj) === '[object Array]'
        }
    }

    if (Array.isArray(props.interests) || isArray) {
        if (props.interests.length < 1) {
            return <div/>
        }
        if ( ! disabled) {
            return (
                <div className="my-divTableRow">
                    <div className="my-divTableCellLeft">&nbsp;</div>
                    <div className="my-divTableCell">
                        <textarea
                            value={textInterests}
                            rows="5"
                            cols="48"
                            name="interests"
                            onChange={e => setTextInterests2(e.target.value)}
                        />
                    </div>
                    <div className="my-divTableCellRight">&nbsp;</div>
                </div>
            )
        }
        return props.interests.map((interest) => (
            <div className="my-divTableRow" key={interest}>
                <div className="my-divTableCellLeft">&nbsp;</div>
                <div className="my-divTableCell">
                    <Input value={interest} disabled={props.disabled}/>
                </div>
                <div className="my-divTableCellRight">&nbsp;</div>
            </div>
        ))
    } else {
        return <div/>
    }
}

UserInterests.propTypes = {
    setInterests: PropTypes.func.isRequired
}