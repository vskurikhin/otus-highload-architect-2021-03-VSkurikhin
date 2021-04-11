
import React, {useEffect} from 'react'
import {compose} from "redux";
import {withRouter} from "react-router-dom";
import {connect} from "react-redux";

import UserDetails from "../UserDefails/UserDetails"
import {getProfileFetch, logoutUser} from "../../redux/currentUser";

export const UserForm = props => {

    const {match} = props

    useEffect(() => props.getProfile(), [])

    return (
        <div className="login-wrapper">
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
                <UserDetails id={match.params.id}/>
            </div>
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
