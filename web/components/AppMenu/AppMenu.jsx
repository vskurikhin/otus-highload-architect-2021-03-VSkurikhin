import {getProfileFetch} from "../../redux/currentUser";

import PropTypes from "prop-types";
import React, {Component} from "react";
import {Link, withRouter} from "react-router-dom";
import {Menu, Container} from "semantic-ui-react";
import {compose} from "redux";
import {connect} from "react-redux";

class AppMenu extends Component {

    disabledUser = (mayBeDisabled, mayBeAdmin) => {
        const {currentUser} = this.props.user
        if (mayBeAdmin) {
            return currentUser.Id === null
                || currentUser.Id === undefined || currentUser.Username !== 'root'
        } if (mayBeDisabled) {
            return currentUser.Id === null
                || currentUser.Id === undefined
        }
        return false
    }

    render() {

        let menuItems = [];
        for (let i = 0; i < this.props.items.length; i++) {
            if (this.props.items[i].length < 2) {
                console.error('HeaderMenu: items format should be ["name", "route"]')
                break;
            }
            const name = this.props.items[i][0]
            const route = this.props.items[i][1]
            const mayBeDisabled = this.props.items[i][2]
            const mayBeAdmin = this.props.items[i][3]
            menuItems.push(
                <Menu.Item
                    key={"item-" + i}
                    disabled={this.disabledUser(mayBeDisabled, mayBeAdmin)}
                    index={i}
                    as={this.disabledUser(mayBeDisabled, mayBeAdmin) ? null : Link}
                    to={route}
                    header={i === 0}
                    active={route === this.props.location.pathname}
                >{name}
                </Menu.Item>
            );
        }

        return (
            <Menu>
                <Container>{menuItems}</Container>
            </Menu>
        )
    }
}

AppMenu.propTypes = {
    onItemClick: PropTypes.func.isRequired,
    items: PropTypes.arrayOf(PropTypes.array.isRequired).isRequired
}

const mapStateToProps = state => ({
    user: state.currentUser
})

const mapDispatchToProps = dispatch => ({
    getProfile: () => dispatch(getProfileFetch())
})

export default compose(
    withRouter,
    connect(mapStateToProps, mapDispatchToProps)
)(AppMenu)
