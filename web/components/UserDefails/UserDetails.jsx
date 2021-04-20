import React, {useEffect, useState} from "react"
import {Dropdown, Input} from "semantic-ui-react"
import {useHistory} from "react-router-dom"

import PropTypes, {string} from "prop-types";
import UserInterests from "../UserInterests/UserInterests";
import {CITY_OPTIONS, SEX_OPTIONS} from "../../consts"
import {POST} from "../../lib/consts";

async function save(value) {
    return fetch('/save', {
        body: JSON.stringify(value),
        ...POST
    }).then(data => data.json())
}

const UserDetails = props => {

    const history = useHistory()
    const [error, setError] = useState(null)
    const [isLoaded, setIsLoaded] = useState(false)
    const [id, setId] = useState(null)
    const [username, setUsername] = useState(null)
    const [name, setName] = useState("")
    const [surName, setSurName] = useState("")
    const [age, setAge] = useState(0)
    const [sex, setSex] = useState("0")
    const [city, setCity] = useState("Murray Hill")
    const [interests, setInterests] = useState([])

    const handleSave = async e => {
        e.preventDefault()
        await save({
            Id: id,
            Username: username,
            Name: name,
            SurName: surName,
            Age: age,
            Sex: sex,
            City: city,
            Interests: interests
        })
    }

    const setItem = result => {
        const {Id, Username, Name, SurName, Age, Sex, City, Interests} = result
        console.debug("Sex")
        console.debug(Sex)
        setId(Id)
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

    const saveForm = () => (
        <div className="my-divTableRow">
            <div className="my-divTableCellLeft">&nbsp;</div>
            <div className="my-divTableCell">
                <button type="button" onClick={handleSave}>Save</button>
            </div>
            <div className="my-divTableCellRight">&nbsp;</div>
        </div>
    )

    const onSexChange = (e, data) => setSex(data.value)

    const onCityChange = (e, data) => setCity(data.value)

    const onCitySearchChange = (e, data) => setCity(data.searchQuery)

    const dropdownSex = () => (
        <Dropdown
            defaultValue={sex}
            item
            fluid
            selection
            onChange={onSexChange}
            options={SEX_OPTIONS}
        />
    )

    const inputDisabledSex = (disabled) => (
        <Input value={sex === "0" ? 'Male' : 'Female'} disabled={true}/>
    )

    useEffect(getItem, [])

    if (error) {
        return <div>Ошибка: {error.message}</div>
    } else if (!isLoaded) {
        return <div>Загрузка...</div>
    } else if (username) {
        const disabled = username !== props.user.currentUser.Username
        props.setIsFriend(id !== props.user.currentUser.Id)
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
                            <Input value={name} disabled={disabled} onChange={e => setName(e.target.value)}/>
                        </div>
                        <div className="my-divTableCellRight">&nbsp;</div>
                    </div>
                    <div className="my-divTableRow">
                        <div className="my-divTableCellLeft">&nbsp;</div>
                        <div className="my-divTableCell">
                            <p className="my-p-label">Surname:</p>
                            <Input value={surName} disabled={disabled} onChange={e => setSurName(e.target.value)}/>
                        </div>
                        <div className="my-divTableCellRight">&nbsp;</div>
                    </div>
                    <div className="my-divTableRow">
                        <div className="my-divTableCellLeft">&nbsp;</div>
                        <div className="my-divTableCell">
                            <p className="my-p-label">Age:</p>
                            <Input value={age} disabled={disabled} onChange={e => setAge(e.target.value)}/>
                        </div>
                        <div className="my-divTableCellRight">&nbsp;</div>
                    </div>
                    <div className="my-divTableRow">
                        <div className="my-divTableCellLeft">&nbsp;</div>
                        <div className="my-divTableCell">
                            <p className="my-p-label">Sex:</p>
                            {disabled ? inputDisabledSex() : dropdownSex() }
                        </div>
                        <div className="my-divTableCellRight">&nbsp;</div>
                    </div>
                    <div className="my-divTableRow">
                        <div className="my-divTableCellLeft">&nbsp;</div>
                        <div className="my-divTableCell">
                            <p className="my-p-label">City</p>
                            <Dropdown
                                disabled={disabled}
                                value={city}
                                fluid
                                search
                                selection
                                onChange={disabled ? null : onCityChange}
                                onSearchChange={disabled ? null : onCitySearchChange}
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
                    <UserInterests disabled={disabled} interests={interests} {...props} />
                    { ! disabled ? saveForm() : <div/>}
                </div>
            )
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


UserDetails.propTypes = {
    setIsFriend: PropTypes.func.isRequired
}
