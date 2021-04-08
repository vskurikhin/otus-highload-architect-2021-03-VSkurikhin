import React, {useEffect, useState} from 'react'
import {Table} from 'semantic-ui-react'

const TableOfUsers = () => {

    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [items, setItems] = useState([]);

    const getResult = result => {
        setIsLoaded(true);
        setItems(result);
    }

    const getError = error => {
        setIsLoaded(true);
        setError(error);
    }

    const getItems = () => {
        fetch("/user")
            .then(res => res.json())
            .then(getResult, getError);
    }

    useEffect(getItems, [])

    const handleClick = e => {
        e.preventDefault();
        const {target} = e
        const {parentElement} = target
        if (parentElement) {
            console.log(parentElement.id)
        }
    };

    if (error) {
        return <div>Ошибка: {error.message}</div>;
    } else if (!isLoaded) {
        return <div>Загрузка...</div>;
    } else {
        return (
            <Table celled selectable>
                <Table.Header>
                    <Table.Row>
                        <Table.HeaderCell>FirstName</Table.HeaderCell>
                        <Table.HeaderCell>SurName</Table.HeaderCell>
                        <Table.HeaderCell>Age</Table.HeaderCell>
                        <Table.HeaderCell>Sex</Table.HeaderCell>
                        <Table.HeaderCell>City</Table.HeaderCell>
                    </Table.Row>
                </Table.Header>

                <Table.Body>
                    {items.map(({Id, Name, SurName, Age, Sex, Interests, City}) => (
                        <Table.Row key={Id} id={Id}>
                            <Table.Cell onClick={handleClick}>{Name}</Table.Cell>
                            <Table.Cell onClick={handleClick}>{SurName}</Table.Cell>
                            <Table.Cell onClick={handleClick}>{Age}</Table.Cell>
                            <Table.Cell onClick={handleClick}>{Sex === 0 ? "Male" : "Female"}</Table.Cell>
                            <Table.Cell onClick={handleClick}>{City}</Table.Cell>
                        </Table.Row>
                    ))}
                </Table.Body>
            </Table>
        );
    }
}

export default TableOfUsers
