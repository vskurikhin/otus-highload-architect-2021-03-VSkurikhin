import './Messages.css'
import TableOfMessages from "../TableOfMessages/TableOfMessages";
import {AFTER_LOGIN, POST} from "../../lib/consts";
import {getProfileFetch} from "../../redux/currentUser";

import React, {useEffect, useState} from 'react'
import {Input} from 'semantic-ui-react'
import {useHistory, withRouter} from "react-router-dom"
import {compose} from "redux";
import {connect} from "react-redux";

async function postMessage(message) {
    return fetch('/message', {
        body: JSON.stringify(message),
        ...POST
    }).then(data => data.json())
}

function Messages(props) {

    const [counter, setCounter] = useState()
    const [message, setMessage] = useState()
    const [username, setUserName] = useState()
    const history = useHistory()

    const getResult = result => {
        if (result.Code > 399 && result.Message) {
            history.push('/error/' + result.Message)
        }
        setCounter(result)
    }

    const getCounter = () => {
        console.debug('props')
        console.debug(props)
        if (props.user.currentUser !== undefined && props.user.currentUser.Username !== undefined) {
            fetch("/counter/" + props.user.currentUser.Username)
                .then(res => res.json())
                .then(getResult)
        }
    }

    useEffect(() => {
        setTimeout(getCounter, 1500);
    }, [props.user.currentUser]); // eslint-disable-line react-hooks/exhaustive-deps

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
                                Total: {counter !== undefined ? (counter.Total !== undefined ? counter.Total : 0) : 0} / Unread: {counter !== undefined ? (counter.Unread !== undefined ? counter.Unread : 0) : 0}
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

const mapStateToProps = state => ({
    user: state.currentUser
})

const mapDispatchToProps = dispatch => ({
    getProfile: () => dispatch(getProfileFetch())
})

export default compose(
    withRouter,
    connect(mapStateToProps, mapDispatchToProps)
)(Messages)
