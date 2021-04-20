import React, {useEffect, useState} from 'react'
import {compose} from "redux";
import {withRouter} from "react-router-dom";
import {connect} from "react-redux";

import UserDetails from "../UserDefails/UserDetails"
import {getProfileFetch, logoutUser} from "../../redux/currentUser";
import {POST} from "../../lib/consts";

async function friend(credentials) {
    return fetch('/friend', {
        body: JSON.stringify(credentials),
        ...POST
    }).then(data => data.json())
}

export const UserForm = props => {

    const [match] = useState(props.match)
    const [currentUser] = useState(props.user.currentUser)
    const [isFriend, setIsFriend] = useState(true)

    const handleSubmit = async e => {
        e.preventDefault()
        await friend({
            UserId: currentUser.Id,
            FriendId: match.params.id
        })
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

    useEffect(() => props.getProfile(), [])

    const notOwn = match.params.id !== props.user.currentUser.Id

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
                    <UserDetails id={match.params.id} setIsFriend={setIsFriend} {...props}/>
                </div>
                {isFriend ? addFriend(props) : notOwn ? okFriend(props) : <div/>}
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
