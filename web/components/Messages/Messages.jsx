import './Messages.css'

import {AFTER_LOGIN, POST} from "../../lib/consts";

import React, {useState} from 'react'
import {Input} from 'semantic-ui-react'
import {useHistory} from "react-router-dom"
import TableOfMessages from "../TableOfMessages/TableOfMessages";

async function postMessage(message) {
    return fetch('/message', {
        body: JSON.stringify(message),
        ...POST
    }).then(data => data.json())
}

export default function Messages(props) {

    const [message, setMessage] = useState()
    const [username, setUserName] = useState()
    const history = useHistory()

    const handleSubmit = async e => {
        e.preventDefault()

        const token = await postMessage({
            Message: message,
            ToUser: username,
        })
        if (token) {
            if (token.Code > 399 && token.Message) {
                history.push('/error/' + token.Message)
            } else {
                history.push(AFTER_LOGIN)
            }
        }
    }

    return (
        <div className="signin-wrapper">
            <form onSubmit={handleSubmit}>
                <div className="my-divTable">
                    <div className="my-divTableBody">
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <h1>Message For</h1>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p className="my-p-label">Username:</p>
                                <Input placeholder='Username...' onChange={e => setUserName(e.target.value)}/>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p className="my-p-label">Message</p>
                                <textarea
                                    rows="5"
                                    cols="48"
                                    name="message"
                                    onChange={e => setMessage(e.target.value)}
                                />
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <button
                                    type="submit"
                                >Submit
                                </button>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                    </div>
                </div>
            </form>
            <div className="login-wrapper">
                <div className="my-divTable">
                    <div className="my-divTableBody">
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <h1>Dialog messages</h1>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <TableOfMessages {...props}/>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}
