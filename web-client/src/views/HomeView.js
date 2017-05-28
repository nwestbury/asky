/**
 * Created by James on 2017-05-13.
 */

import React, { Component } from 'react';
import TabBar from '../components/tabbar/TabBar';

const tabItems = [
    { label: "Tab 1", link: "/home", accentColor: "#63b6e5" },
    { label: "Tab 2", link: "/home/tab2", accentColor: "#63b6e5" },
];

const headerStyle = {
    marginBottom: 40,
    width: "100%",
};

class HomeView extends Component {
    render() {
        return (
            <div style={headerStyle}>
                <h2 className="align-top" style={{display: "inline-block"}}>Home Page</h2>
                <TabBar items={tabItems}/>
            </div>
        );
    }
}

export default HomeView;