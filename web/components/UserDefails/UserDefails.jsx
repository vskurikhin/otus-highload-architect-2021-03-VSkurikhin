import React, {useEffect, useState} from "react";
import {Dropdown, Input, Table} from "semantic-ui-react";
import {CITY_OPTIONS, SEX_OPTIONS} from "../../consts";
import {useHistory} from "react-router-dom";

const UserDefails = (props) => {

    const history = useHistory();
    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [item, setItem] = useState({});

    const getResult = result => {
        setIsLoaded(true);
        if (result.code && result.message) {
            throw {
                code: result.code,
                message: result.message
            }
        }
        setItem(result);
    }

    const getError = error => {
        setIsLoaded(true);
        setError(error);
    }

    const getItem = () => {
        fetch("/user/" + props.id)
            .then(res => res.json())
            .then(getResult, getError);
    }

    useEffect(getItem, [])

    const handleClick = e => {
        e.preventDefault();
        const {target} = e
        const {parentElement} = target
        if (parentElement) {
            console.log(parentElement.id)
        }
    };

    console.log(item);

    if (error) {
        return <div>Ошибка: {error.message}</div>;
    } else if (!isLoaded) {
        return <div>Загрузка...</div>;
    } else {
        try {
            return (
                <div className="my-divTableBody">
                    <div className="my-divTableRow">
                        <div className="my-divTableCellLeft">&nbsp;</div>
                        <div className="my-divTableCell">
                            <p className="my-p-label">Username:</p>
                            <Input value={item.Username} disabled={true}/>
                        </div>
                        <div className="my-divTableCellRight">&nbsp;</div>
                    </div>
                    <div className="my-divTableRow">
                        <div className="my-divTableCellLeft">&nbsp;</div>
                        <div className="my-divTableCell">
                            <p className="my-p-label">Firstname:</p>
                            <Input value={item.Name} disabled={true}/>
                        </div>
                        <div className="my-divTableCellRight">&nbsp;</div>
                    </div>
                    <div className="my-divTableRow">
                        <div className="my-divTableCellLeft">&nbsp;</div>
                        <div className="my-divTableCell">
                            <p className="my-p-label">Surname:</p>
                            <Input value={item.SurName} disabled={true}/>
                        </div>
                        <div className="my-divTableCellRight">&nbsp;</div>
                    </div>
                    <div className="my-divTableRow">
                        <div className="my-divTableCellLeft">&nbsp;</div>
                        <div className="my-divTableCell">
                            <p className="my-p-label">Age:</p>
                            <Input value={item.Age} disabled={true}/>
                        </div>
                        <div className="my-divTableCellRight">&nbsp;</div>
                    </div>
                    <div className="my-divTableRow">
                        <div className="my-divTableCellLeft">&nbsp;</div>
                        <div className="my-divTableCell">
                            <p className="my-p-label">Sex:</p>
                            <Input value={item.Sex === 0 ? 'Male' : 'Female'} disabled={true}/>
                        </div>
                        <div className="my-divTableCellRight">&nbsp;</div>
                    </div>
                    <div className="my-divTableRow">
                        <div className="my-divTableCellLeft">&nbsp;</div>
                        <div className="my-divTableCell">
                            <p className="my-p-label">City</p>
                            <Dropdown
                                disabled={true}
                                defaultValue={item.City}
                                value={item.City}
                                fluid
                                search
                                selection
                                options={CITY_OPTIONS}
                            />
                        </div>
                        <div className="my-divTableCellRight">&nbsp;</div>
                    </div>
                    {item && item !== '{}' && Array.isArray(item.Interests) ? (
                        <div className="my-divTableRow">
                            <div className="my-divTableCellLeft">&nbsp;</div>
                            <div className="my-divTableCell">
                                <p className="my-p-label">Interests:</p>
                                {item.Interests.map((interest) => (
                                    <Input value={interest} disabled={true}/>
                                ))}
                            </div>
                            <div className="my-divTableCellRight">&nbsp;</div>
                        </div>
                    ) : <div/>}
                </div>
            );
        } catch (e) {
            console.debug(e);
            history.push('/login');
            return <div/>;
        }
    }
}

export default UserDefails;
