import React, {useEffect, useState} from 'react'
import {Table} from 'semantic-ui-react'
import {useHistory} from "react-router-dom"

export default function TableOfNews() {

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

    const getItems = () => {
        fetch("/news/range/0/99")
            .then(res => res.json())
            .then(getResult, getError)
    }

    const handleClick = e => {
        e.preventDefault()
        const {target} = e
        const {parentElement} = target
        if (parentElement) {
            history.push('/userform/' + parentElement.id)
        }
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
                    <Table.HeaderCell>Title</Table.HeaderCell>
                    <Table.HeaderCell>Content</Table.HeaderCell>
                    <Table.HeaderCell>Public At</Table.HeaderCell>
                </Table.Row>
            </Table.Header>

            <Table.Body>
                {items.map(({id, Title, Content, PublicAt}) => (
                    <Table.Row key={id} id={id}>
                        <Table.Cell onClick={handleClick}>{Title}</Table.Cell>
                        <Table.Cell onClick={handleClick}>{Content}</Table.Cell>
                        <Table.Cell onClick={handleClick}>{PublicAt}</Table.Cell>
                    </Table.Row>
                ))}
            </Table.Body>
        </Table>
    )
}
