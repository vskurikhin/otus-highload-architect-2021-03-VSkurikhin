import UserDetails from "../UserDefails/UserDetails"
import {AFTER_LOGIN, POST} from "../../lib/consts";
import {getProfileFetch, logoutUser} from "../../redux/currentUser";

import React, {useEffect, useState} from 'react'
import {compose} from "redux";
import {connect} from "react-redux";
import {useHistory, withRouter} from "react-router-dom";

async function friend(credentials) {
    return fetch('/friend', {
        body: JSON.stringify(credentials),
        ...POST
    }).then(data => data.json())
}

const UserForm = props => {

    const [error, setError] = useState(null)
    const [isLoaded, setIsLoaded] = useState(false)
    const [item, setItem] = useState({Friend: false})
    const history = useHistory()
    const {match} = props

    const handleSubmit = async e => {
        e.preventDefault()
        await friend({
            UserId: props.user.currentUser.Id,
            FriendId: match.params.id
        })
        history.push(AFTER_LOGIN)
    }

    const addFriend = props => {
        return (
            <div className="my-divTable">
                <div className="my-divTableRow">
                    <div className="my-divTableCellLeft">&nbsp;</div>
                    <div className="my-divTableCell">
                        <button type="submit">Add Friend</button>
                    </div>
                    <div className="my-divTableCellRight">&nbsp;</div>
                </div>
            </div>
        )
    }

    const okFriend = props => {
        return (
            <div className="my-divTable">
                <div className="my-divTableRow">
                    <div className="my-divTableCellLeft">&nbsp;</div>
                    <div className="my-divTableCell">
                        <button type="submit" disabled={true}>Is friend</button>
                    </div>
                    <div className="my-divTableCellRight">&nbsp;</div>
                </div>
            </div>
        )
    }

    const getError = error => {
        setIsLoaded(true)
        setError(error)
    }

    const getItem = async () => {
        await fetch("/user/" + match.params.id)
            .then(res => res.json())
            .then(getResult, getError)
    }

    const getResult = result => {
        setIsLoaded(true)
        if (result.Code > 399 && result.Message) {
            history.push('/error/' + result.Message)
        }
        setItem(result)
    }

    // useEffect(setCurrentUser, [props.user.currentUser])
    useEffect(() => props.getProfile(), [])
    useEffect(getItem, [])

    const notOwn = props.user.currentUser.Id !== match.params.id

    if (error) {
        return <div>Ошибка: {error.message}</div>
    } else if (!isLoaded) {
        return <div>Загрузка...</div>
    }
    const disabled = props.user.currentUser.Id !== item.Id
    return (
        <div className="login-wrapper">
            <form onSubmit={handleSubmit}>
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
                    <UserDetails disabled={disabled} item={item} {...props}/>
                </div>
                { ! item.Friend && notOwn ? addFriend(props) : (notOwn ? okFriend(props) : <div/>)}
            </form>
        </div>
    )
}

const mapStateToProps = state => ({
    user: state.currentUser
})

const mapDispatchToProps = dispatch => ({
    getProfile: () => dispatch(getProfileFetch()),
    logoutUser: () => dispatch(logoutUser())
})

export default compose(
    withRouter,
    connect(mapStateToProps, mapDispatchToProps)
)(UserForm)
