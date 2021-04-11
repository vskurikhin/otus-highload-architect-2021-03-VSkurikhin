
import TableOfUsers from "../TableOfUsers/TableOfUsers"
import {getProfileFetch, logoutUser} from "../../redux/currentUser"

import React, {useEffect} from 'react'
import {compose} from "redux"
import {connect} from "react-redux"
import {useHistory, withRouter} from "react-router-dom"

export const UserList = props => {

    const history = useHistory()

    useEffect(() => props.getProfile(), [])

    try {
        return (
            <div className="login-wrapper">
                <div className="my-divTable">
                    <div className="my-divTableBody">
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <h1>User list</h1>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <TableOfUsers {...props}/>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                    </div>
                </div>
            </div>
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
)(UserList)
