import React, { Component } from 'react';
import { Provider } from 'react-redux';
import {
    BrowserRouter as Router,
    Route,
} from 'react-router-dom'
import store from './store';

import SideBarWrapper, { defaultSidebarItems } from './components/sidebar/SideBar';

import MainPage from './views/MainView';
import HomePage from './views/HomeView';
import LoginPage from './views/LoginView';

const App = () => (
    <Provider store={store}>
        <Router>
            <div>
                <SideBarWrapper items={defaultSidebarItems}>
                    <Route exact path="/" component={MainPage}/>
                    <Route path="/home" component={HomePage} />
                    <Route path="/login" component={LoginPage}/>
                </SideBarWrapper>
            </div>
        </Router>
    </Provider>
);

export default App;
