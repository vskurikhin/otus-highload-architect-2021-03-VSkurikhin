import React, {useEffect, useState} from "react"
import {Dropdown, Input} from "semantic-ui-react"
import {useHistory} from "react-router-dom"

import {CITY_OPTIONS} from "../../consts"
import UserInterests from "../UserInterests/UserInterests";

const UserDetails = props => {

    const history = useHistory()
    const [error, setError] = useState(null)
    const [isLoaded, setIsLoaded] = useState(false)
    const [username, setUsername] = useState(null)
    const [name, setName] = useState("")
    const [surName, setSurName] = useState("")
    const [age, setAge] = useState(0)
    const [sex, setSex] = useState(0)
    const [city, setCity] = useState("Murray Hill")
    const [interests, setInterests] = useState([])

    const setItem = result => {
        const {Username, Name, SurName, Age, Sex, City, Interests} = result
        setUsername(Username)
        setName(Name)
        setSurName(SurName)
        setAge(Age)
        setSex(Sex)
        setCity(City)
        setInterests(Interests)
        setIsLoaded(true)
    }

    const getResult = result => {
        setIsLoaded(true)
        if (result.code && result.message) {
            throw {
                code: result.code,
                message: result.message
            }
        }
        setItem(result)
    }

    const getError = error => {
        setIsLoaded(true)
        setError(error)
    }

    const getItem = async () => {
        await fetch("/user/" + props.id)
            .then(res => res.json())
            .then(getResult, getError)
    }

    useEffect(getItem, [])

    if (error) {
        return <div>Ошибка: {error.message}</div>
    } else if (!isLoaded) {
        return <div>Загрузка...</div>
    } else if (username) {
        try {
            return (
                <div className="my-divTableBody">
                    <div className="my-divTableRow">
                        <div className="my-divTableCellLeft">&nbsp;</div>
                        <div className="my-divTableCell">
                            <p className="my-p-label">Username:</p>
                            <Input value={username} disabled={true}/>
                        </div>
                        <div className="my-divTableCellRight">&nbsp;</div>
                    </div>
                    <div className="my-divTableRow">
                        <div className="my-divTableCellLeft">&nbsp;</div>
                        <div className="my-divTableCell">
                            <p className="my-p-label">Firstname:</p>
                            <Input value={name} disabled={true}/>
                        </div>
                        <div className="my-divTableCellRight">&nbsp;</div>
                    </div>
                    <div className="my-divTableRow">
                        <div className="my-divTableCellLeft">&nbsp;</div>
                        <div className="my-divTableCell">
                            <p className="my-p-label">Surname:</p>
                            <Input value={surName} disabled={true}/>
                        </div>
                        <div className="my-divTableCellRight">&nbsp;</div>
                    </div>
                    <div className="my-divTableRow">
                        <div className="my-divTableCellLeft">&nbsp;</div>
                        <div className="my-divTableCell">
                            <p className="my-p-label">Age:</p>
                            <Input value={age} disabled={true}/>
                        </div>
                        <div className="my-divTableCellRight">&nbsp;</div>
                    </div>
                    <div className="my-divTableRow">
                        <div className="my-divTableCellLeft">&nbsp;</div>
                        <div className="my-divTableCell">
                            <p className="my-p-label">Sex:</p>
                            <Input value={sex === 0 ? 'Male' : 'Female'} disabled={true}/>
                        </div>
                        <div className="my-divTableCellRight">&nbsp;</div>
                    </div>
                    <div className="my-divTableRow">
                        <div className="my-divTableCellLeft">&nbsp;</div>
                        <div className="my-divTableCell">
                            <p className="my-p-label">City</p>
                            <Dropdown
                                disabled={true}
                                value={city}
                                fluid
                                search
                                selection
                                options={CITY_OPTIONS}
                            />
                        </div>
                        <div className="my-divTableCellRight">&nbsp;</div>
                    </div>
                    <div className="my-divTableRow">
                        <div className="my-divTableCellLeft">&nbsp;</div>
                        <div className="my-divTableCell">
                            <p className="my-p-label">Interests:</p>
                        </div>
                        <div className="my-divTableCellRight">&nbsp;</div>
                    </div>
                    <UserInterests interests={interests} {...props} />
                </div>)
        } catch (e) {
            console.debug(e)
            history.push('/login')
            return <div/>
        }
    } else {
        return <div/>
    }
}

export default UserDetails
