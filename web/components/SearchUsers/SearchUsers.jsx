import React, {useEffect, useState} from 'react'
import {Input, Table} from 'semantic-ui-react'
import {useHistory} from "react-router-dom"

export default function SearchUsers() {

    const [error, setError] = useState(null)
    const [isLoaded, setIsLoaded] = useState(false)
    const [items, setItems] = useState([])
    const [name, setName] = useState("")
    const [surname, setSurName] = useState("")
    const history = useHistory()

    const getResult = result => {
        setIsLoaded(true)
        if (result.Code > 399 && result.Message) {
            history.push('/error/' + result.Message)
        }
        setItems(result)
    }

    const getError = error => {
        setIsLoaded(true)
        setError(error)
    }

    const getItems = () => {
        if (name && name.length > 0 && surname && surname.length > 0) {
            fetch("/users/search/" + name + "/" + surname)
                .then(res => res.json())
                .then(getResult, getError)
        } else if (name && name.length > 0) {
            fetch("/users/search-by/name/" + name)
                .then(res => res.json())
                .then(getResult, getError)
        } else if (surname && surname.length > 0) {
            fetch("/users/search-by/surname/" + surname)
                .then(res => res.json())
                .then(getResult, getError)
        }
    }

    const handleClick = e => {
        e.preventDefault()
        const {target} = e
        const {parentElement} = target
        if (parentElement) {
            history.push('/userform/' + parentElement.id)
        }
    }

    const getTable = () => {
        if (error) {
            return <div>Ошибка: {error.message}</div>
        } else if (!isLoaded) {
            return <div>Загрузка...</div>
        }
        return (
                <Table celled selectable>
                    <Table.Header>
                        <Table.Row>
                            <Table.HeaderCell>FirstName</Table.HeaderCell>
                            <Table.HeaderCell>SurName</Table.HeaderCell>
                            <Table.HeaderCell>City</Table.HeaderCell>
                            <Table.HeaderCell>Friend</Table.HeaderCell>
                        </Table.Row>
                    </Table.Header>

                    <Table.Body>
                        {items.map(({Id, Name, SurName, Age, Sex, Interests, City, Friend}) => (
                            <Table.Row key={Id} id={Id}>
                                <Table.Cell onClick={handleClick}>{Name}</Table.Cell>
                                <Table.Cell onClick={handleClick}>{SurName}</Table.Cell>
                                <Table.Cell onClick={handleClick}>{City}</Table.Cell>
                                <Table.Cell onClick={handleClick}>{Friend ? "✔" : ""}</Table.Cell>
                            </Table.Row>
                        ))}
                    </Table.Body>
                </Table>
        )
    }

    useEffect(() => {
        setTimeout(getItems, 1500);
    }, [name, surname]); // eslint-disable-line react-hooks/exhaustive-deps

    return (
        <div>
            <div className="my-divTable">
                <div className="my-divTableBody">
                    <div className="my-divTableRow">
                        <div className="my-divTableCellLeft">&nbsp;</div>
                        <div className="my-divTableCell">
                            <Input value={name || ''} onChange={e => setName(e.target.value)}/>
                            <Input value={surname || ''} onChange={e => setSurName(e.target.value)}/>
                        </div>
                        <div className="my-divTableCellRight">&nbsp;</div>
                    </div>
                </div>
            </div>
            {getTable()}
        </div>
    )
}
