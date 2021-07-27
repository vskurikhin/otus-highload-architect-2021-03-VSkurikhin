import React, {useEffect, useState} from 'react'
import {Table} from 'semantic-ui-react'
import {useHistory} from "react-router-dom"

export default function TableOfMessages(props) {

    const [error, setError] = useState(null)
    const [isLoaded, setIsLoaded] = useState(false)
    const [items, setItems] = useState([])
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

    const getItems = userId => {
        console.debug('props')
        console.debug(props)
        fetch("/messages")
            .then(res => res.json())
            .then(getResult, getError)
    }

    useEffect(getItems, [])

    if (error) {
        return <div>Ошибка: {error.message}</div>
    } else if (!isLoaded) {
        return <div>Загрузка...</div>
    }
    return (
        <Table celled selectable>
            <Table.Header>
                <Table.Row>
                    <Table.HeaderCell>FromUser</Table.HeaderCell>
                    <Table.HeaderCell>ToUser</Table.HeaderCell>
                    <Table.HeaderCell>Message</Table.HeaderCell>
                </Table.Row>
            </Table.Header>

            <Table.Body>
                {items.map(({Id, FromUser, ToUser, Message}) => (
                    <Table.Row key={Id} id={Id}>
                        <Table.Cell>{FromUser}</Table.Cell>
                        <Table.Cell>{ToUser}</Table.Cell>
                        <Table.Cell>{Message}</Table.Cell>
                    </Table.Row>
                ))}
            </Table.Body>
        </Table>
    )
}
