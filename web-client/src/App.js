import React, { Component } from 'react';
import {
    BrowserRouter as Router,
    Route,
} from 'react-router-dom'

import SideBarWrapper, { defaultSidebarItems } from './components/sidebar/SideBar';

import MainPage from './views/MainView';
import HomePage from './views/HomeView';
import LoginPage from './views/LoginView';

class App extends Component {
  render() {
    return (
        <Router>
            <div>
                <SideBarWrapper items={defaultSidebarItems}>
                    <Route exact path="/" component={MainPage}/>
                    <Route path="/home" component={HomePage} />
                    <Route path="/login" component={LoginPage}/>
                </SideBarWrapper>
            </div>
        </Router>
    );
  }
}

export default App;
