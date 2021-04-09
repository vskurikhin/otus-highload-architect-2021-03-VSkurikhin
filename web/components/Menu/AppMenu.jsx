import React, {useState} from "react";
import {useHistory} from "react-router-dom";
import {Menu} from "semantic-ui-react";

function AppMenu(props) {

    const [activeItem, setActiveItem] = useState();
    const history = useHistory();
    const handleItemClick = (e, { name }) => {
        setActiveItem(name);
        history.push(name);
    }

    return (
        <div className="wrapper">
            <Menu>
                <Menu.Item
                    name='/login'
                    active={activeItem === '/login'}
                    onClick={handleItemClick}
                >Login
                </Menu.Item>
                <Menu.Item
                    name='/signin'
                    active={activeItem === '/signin'}
                    onClick={handleItemClick}
                >Sign-in
                </Menu.Item>
                <Menu.Item
                    disabled={props.disabledUserList}
                    name='/userlist'
                    active={activeItem === '/userlist'}
                    onClick={handleItemClick}
                >User list
                </Menu.Item>
            </Menu>
        </div>
    );
}

export default AppMenu;