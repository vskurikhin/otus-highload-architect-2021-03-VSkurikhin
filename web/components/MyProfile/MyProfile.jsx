import UserDetails from "../UserDefails/UserDetails";
import {getProfileFetch} from "../../redux/currentUser";

import React, {useEffect, useState} from "react";
import {compose} from "redux";
import {connect} from "react-redux";
import {useHistory, withRouter} from "react-router-dom";

export const MyProfile = props => {

    const [error, setError] = useState(null)
    const [isLoaded, setIsLoaded] = useState(false)
    const [item, setItem] = useState({Friend: false})
    const history = useHistory()

    const getError = error => {
        setIsLoaded(true)
        setError(error)
    }

    const getItem = async () => {
        await fetch("/user/" + props.user.currentUser.Id)
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

    useEffect(() => props.getProfile(), [])
    useEffect(getItem, [])

    if (error) {
        return <div>Ошибка: {error.message}</div>
    } else if (!isLoaded) {
        return <div>Загрузка...</div>
    }

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
                <UserDetails disabled={false} item={item} {...props}/>
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
)(MyProfile)
