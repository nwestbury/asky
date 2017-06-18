/**
 * Created by James on 2017-05-28.
 */

import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { NavLink } from 'react-router-dom';

import styles from "./tabbar.scss"

class TabBar extends Component {

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
            <div className={styles.tabbar}>
                { items.map((item, index) => (
                    <NavLink
                        activeClassName={styles.selected_label}
                        className={[styles.item, styles.quiet_label, styles.strong].join(" ")}
                        activeStyle={{borderBottom: "solid thin ".concat(item.accentColor)}}
                        style={{borderBottom: "solid thin #eeeeee"}}
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

export default TabBar;