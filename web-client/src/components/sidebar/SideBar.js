/**
 * Created by James on 2017-05-22.
 */

import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { NavLink } from 'react-router-dom';

import "./SideBar.scss"

//TODO: Add icons
export const defaultSidebarItems = [
    { label: "Home", link: "/home", accentColor: "#63b6e5" },
    { label: "Login", link: "/login", accentColor: "#3887BE" }
];

class SideBar extends Component {

    static propTypes = {
        items: PropTypes.arrayOf(
            PropTypes.shape({
                label: PropTypes.string.isRequired,
                link: PropTypes.string.isRequired,
                accentColor: PropTypes.string.isRequired,
            }).isRequired
        ).isRequired,
    };

    render() {
        const { items } = this.props;
        return (
            <div className="sidebar fixed-left dark">
                { items.map((item, index) => (
                    <NavLink
                        activeClassName="sidebar-selected-label"
                        className="sidebar-item clean-link sidebar-quiet-label strong"
                        activeStyle={{borderLeft: "solid medium ".concat(item.accentColor)}}
                        style={{borderLeft: "solid medium #222222"}}
                        key={item.label.toLowerCase()}
                        to={item.link}
                        exact={false}>

                        {item.label}

                    </NavLink>
                )) }
            </div>
        );
    }
}

const SideBarWrapper = (props) => (
  <div style={ { margin: 0, padding: 0 } }>
      <SideBar items={props.items}/>
      <div className="sidebar-content">
          { props.children }
      </div>
  </div>
);

export default SideBarWrapper;