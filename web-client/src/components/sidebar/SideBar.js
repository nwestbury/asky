/**
 * Created by James on 2017-05-22.
 */

import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { NavLink } from 'react-router-dom';

import "./SideBar.css"

//TODO: Add icons
export const defaultSidebarItems = [
    { label: "Home", link:"/" },
    { label: "Login", link:"/login" }
];

class SideBar extends Component {

    render() {
        const { items } = this.props;
        return (
            <div className="sidebar fixed-left dark">
                { items.map((item, index) => (
                    <NavLink
                        activeClassName="selected-label"
                        className="sidebar-item clean-link quiet-label strong"
                        key={item.label.toLowerCase()}
                        to={item.link}
                        exact={true}>

                        {item.label}

                    </NavLink>
                )) }
            </div>
        );
    }
}

SideBar.propTypes = {
    items: PropTypes.objectOf({
        label: PropTypes.string.isRequired,
        link: PropTypes.string.isRequired,
    }).isRequired,
};

const SideBarWrapper = (props) => (
  <div style={ { margin: 0, padding: 0 } }>
      <SideBar items={props.items}/>
      <div className="content">
          { props.children }
      </div>
  </div>
);

export default SideBarWrapper;