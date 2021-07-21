
import Modal from "../Modal/Modal";

import React, {useState} from "react";
import {useHistory} from "react-router-dom";

export default function Error(props) {

    const history = useHistory()
    const [match] = useState(props.match)

    const hideModal = () => {
        history.push('/login')
    }

    const {params} = match

    console.debug("Error")
    console.debug(params.msg)

    return (
        <Modal show={true} handleClose={hideModal} {...props}>
            <p>Modal</p>
        </Modal>
    )
}
