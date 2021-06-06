import React, {useEffect, useRef, useState} from 'react'
import {Table} from 'semantic-ui-react'
import {useHistory} from "react-router-dom"

const FETCH = {Method: "fetch", Offset: 0, Limit: 99}

export default function TableOfNews() {

    const [error, setError] = useState(null)
    const [isLoaded, setIsLoaded] = useState(false)
    const [items, setItems] = useState([])
    const [newMessage] = useState(FETCH);

    const history = useHistory()
    const socket = useRef(null)

    useEffect(() => {
        socket.current = new WebSocket("ws://localhost:8080/ws-newslist");
        socket.current.onopen = () => {
            console.debug("ws opened")
            socket.current.send(JSON.stringify(newMessage))
        }
        socket.current.onclose = () => console.debug("ws closed")
        socket.current.onmessage = (msg) => {
            getResult(JSON.parse(msg.data))
        }
        return () => {
            socket.current.close();
        };
    }, []);

    const getResult = result => {
        setIsLoaded(true)
        if (result.Code > 399 && result.Message) {
            history.push('/error/' + result.Message)
        } else if (result.Code === 1 && result.Message === "push") {
            socket.current.send(JSON.stringify(newMessage))
        } else {
            setItems(result)
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
                {items.map(({Id, Title, Content, PublicAt}) => (
                    <Table.Row key={Id} id={Id}>
                        <Table.Cell key={Id + ".Title"} onClick={handleClick}>{Title}</Table.Cell>
                        <Table.Cell key={Id + ".Content"} onClick={handleClick}>{Content}</Table.Cell>
                        <Table.Cell key={Id + ".PublicAt"} onClick={handleClick}>{PublicAt}</Table.Cell>
                    </Table.Row>
                ))}
            </Table.Body>
        </Table>
    )
}
