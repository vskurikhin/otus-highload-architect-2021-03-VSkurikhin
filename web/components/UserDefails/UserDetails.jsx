import UserInterests from "../UserInterests/UserInterests";
import {CITY_OPTIONS, SEX_OPTIONS} from "../../consts"
import {POST} from "../../lib/consts";

import React, {useEffect, useState} from "react"
import {Dropdown, Input} from "semantic-ui-react"

async function save(value) {
    return fetch('/save', {
        body: JSON.stringify(value),
        ...POST
    }).then(data => data.json())
}

export default function UserDetails(props) {

    const [age, setAge] = useState("")
    const [city, setCity] = useState("")
    const [disabled, setDisabled] = useState("")
    const [id, setId] = useState("")
    const [interests, setInterests] = useState("")
    const [name, setName] = useState("")
    const [sex, setSex] = useState("")
    const [surName, setSurName] = useState("")
    const [username, setUsername] = useState("")

    const handleSave = async e => {
        e.preventDefault()
        await save({
            Id: id,
            Username: username,
            Name: name,
            SurName: surName,
            Age: parseInt(age),
            Sex: sex === "1" ? 1 : 0,
            City: city,
            Interests: interests
        })
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
            defaultValue={sex || '0'}
            value={sex || '1'}
            item
            fluid
            selection
            onChange={onSexChange}
            options={SEX_OPTIONS}
        />
    )

    const inputDisabledSex = () => (
        <Input value={sex === "1" ? 'Female' : 'Male'} disabled={true}/>
    )

    function setItem() {
        setId(props.item.Id)
        setUsername(props.item.Username)
        setName(props.item.Name)
        setSurName(props.item.SurName)
        setAge(props.item.Age)
        setSex(props.item.Sex === 1 ? "1" : "0")
        setCity(props.item.City)
        setInterests(props.item.Interests)
    }

    useEffect(() => setItem(props.item), [props.item])
    useEffect(() => setDisabled(props.disabled), [props.disabled])

    return (
        <div className="my-divTableBody">
            <div className="my-divTableRow">
                <div className="my-divTableCellLeft">&nbsp;</div>
                <div className="my-divTableCell">
                    <p className="my-p-label">Username:</p>
                    <Input value={username || ''} disabled={true}/>
                </div>
                <div className="my-divTableCellRight">&nbsp;</div>
            </div>
            <div className="my-divTableRow">
                <div className="my-divTableCellLeft">&nbsp;</div>
                <div className="my-divTableCell">
                    <p className="my-p-label">Firstname:</p>
                    <Input value={name || ''} disabled={!!disabled} onChange={e => setName(e.target.value)}/>
                </div>
                <div className="my-divTableCellRight">&nbsp;</div>
            </div>
            <div className="my-divTableRow">
                <div className="my-divTableCellLeft">&nbsp;</div>
                <div className="my-divTableCell">
                    <p className="my-p-label">Surname:</p>
                    <Input value={surName || ''} disabled={!!disabled} onChange={e => setSurName(e.target.value)}/>
                </div>
                <div className="my-divTableCellRight">&nbsp;</div>
            </div>
            <div className="my-divTableRow">
                <div className="my-divTableCellLeft">&nbsp;</div>
                <div className="my-divTableCell">
                    <p className="my-p-label">Age:</p>
                    <Input value={age || ''} disabled={!!disabled} onChange={e => setAge(e.target.value)}/>
                </div>
                <div className="my-divTableCellRight">&nbsp;</div>
            </div>
            <div className="my-divTableRow">
                <div className="my-divTableCellLeft">&nbsp;</div>
                <div className="my-divTableCell">
                    <p className="my-p-label">Sex:</p>
                    {!!disabled ? inputDisabledSex() : dropdownSex()}
                </div>
                <div className="my-divTableCellRight">&nbsp;</div>
            </div>
            <div className="my-divTableRow">
                <div className="my-divTableCellLeft">&nbsp;</div>
                <div className="my-divTableCell">
                    <p className="my-p-label">City</p>
                    <Dropdown
                        disabled={!!disabled}
                        value={city || ''}
                        fluid
                        search
                        selection
                        onChange={!!disabled ? null : onCityChange}
                        onSearchChange={!!disabled ? null : onCitySearchChange}
                        options={CITY_OPTIONS}
                    />
                </div>
                <div className="my-divTableCellRight">&nbsp;</div>
            </div>
            <div className="my-divTableRow">
                <div className="my-divTableCellLeft">&nbsp;</div>
                <div className="my-divTableCell">
                    <p className="my-p-label">Interests</p>
                </div>
                <div className="my-divTableCellRight">&nbsp;</div>
            </div>
            <UserInterests disabled={!!disabled} interests={interests} setInterests={setInterests} {...props} />
            { ! !!disabled ? saveForm() : <div/>}
        </div>
    )
}
