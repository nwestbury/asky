/**
 * Created by James on 2017-05-22.
 */

import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { NavLink } from 'react-router-dom';

import sidebarStyles from "./SideBar.scss";

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
            <div className={[sidebarStyles.sidebar, sidebarStyles.fixed_left, sidebarStyles.dark].join(" ")}>
                { items.map((item, index) => (
                    <NavLink
                        activeClassName={sidebarStyles.selected_label}
                        className={[sidebarStyles.item, sidebarStyles.clean_link, sidebarStyles.quiet_label, sidebarStyles.strong].join(" ")}
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
      <div className={sidebarStyles.content}>
          { props.children }
      </div>
  </div>
);

export default SideBarWrapper;