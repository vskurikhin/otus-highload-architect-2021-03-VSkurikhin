
import './Signin.css'

import PropTypes from 'prop-types'
import React, {useState} from 'react'
import {Dropdown, Input} from 'semantic-ui-react'
import {useHistory} from "react-router-dom"

import {CITY_OPTIONS, SEX_OPTIONS} from "../../consts"
import {POST} from "../../lib/consts";

async function signinUser(credentials) {
    return fetch('/signin', {
        body: JSON.stringify(credentials),
        ...POST
    }).then(data => data.json())
}

export default function Signin({setToken}) {

    const history = useHistory()
    const [username, setUserName] = useState()
    const [password, setPassword] = useState()
    const [name, setName] = useState()
    const [surname, setSurname] = useState()
    const [age, setAge] = useState("41")
    const [sex, setSex] = useState("0")
    const [city, setCity] = useState("Москва")
    const [interests, setInterests] = useState()

    const handleSubmit = async e => {
        e.preventDefault()
        const token = await signinUser({
            username,
            password,
            name,
            surname,
            age,
            sex,
            city,
            interests
        })
        setToken(token)
        if (token){
            history.push('/userlist')
        }
    }

    const onSexChange = (e, data) => setSex(data.value)

    const onCityChange = (e, data) => setCity(data.value)

    const onCitySearchChange = (e, data) => setCity(data.searchQuery)

    return (
        <div className="signin-wrapper">
            <form onSubmit={handleSubmit}>
                <div className="my-divTable">
                    <div className="my-divTableBody">
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <h1>For register Sign in please</h1>
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
                                <p className="my-p-label">Password:</p>
                                <Input type="password" name="password" placeholder='password...' onChange={e => setPassword(e.target.value)}/>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p className="my-p-label">Firstname:</p>
                                <Input type="text" name="name" placeholder='Name...' onChange={e => setName(e.target.value)}/>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p className="my-p-label">Surname:</p>
                                <Input type="text" name="surname" placeholder='Surname...' onChange={e => setSurname(e.target.value)}/>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p className="my-p-label">Age:</p>
                                <Input type="text" name="surname" value={age} onChange={e => setAge(e.target.value)}/>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p className="my-p-label">Sex:</p>
                                <Dropdown
                                    defaultValue={sex}
                                    item
                                    fluid
                                    selection
                                    onChange={onSexChange}
                                    options={SEX_OPTIONS}
                                />
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p className="my-p-label">City</p>
                                <Dropdown
                                    defaultValue={city}
                                    fluid
                                    search
                                    selection
                                    onChange={onCityChange}
                                    onSearchChange={onCitySearchChange}
                                    options={CITY_OPTIONS}
                                />
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p className="my-p-label">Interests</p>
                                <textarea
                                    rows="5"
                                    cols="48"
                                    name="interests"
                                    onChange={e => setInterests(e.target.value)}
                                />
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <button type="submit">Submit</button>
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                    </div>
                </div>
            </form>
        </div>
    )
}

Signin.propTypes = {
    setToken: PropTypes.func.isRequired
}
