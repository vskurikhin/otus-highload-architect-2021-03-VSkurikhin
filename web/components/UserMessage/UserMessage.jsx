
import Messages from "../Messages/Messages";
import {getProfileFetch, logoutUser} from "../../redux/currentUser"

import React, {useEffect} from 'react'
import {compose} from "redux"
import {connect} from "react-redux"
import {useHistory, withRouter} from "react-router-dom"

const UserMessage = props => {

    const history = useHistory()

    useEffect(() => props.getProfile(), [])

    try {
        return (
            <Messages {...props}/>
        )
    } catch (e) {
        console.debug(e)
        history.push('/login')
        return <div/>
    }
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
)(UserMessage)
