import './modal.css'

import React, {useState} from "react";
import {useHistory} from "react-router-dom";

export default function Modal(props) {

    const history = useHistory()
    const [match] = useState(props.match)

    const handleClose = () => {
        history.push('/login')
    }

    const {params} = match

    return (
        <div id="openModal" className="modalDialog">
            <div>
                <a href="#close" title="Close" className="close" onClick={handleClose}>>X</a>
                <h2>Error</h2>
                <p>{params.msg}</p>
            </div>
        </div>
    )
}
