import React, { Component } from 'react';
import {
    BrowserRouter as Router,
    Route,
    NavLink,
} from 'react-router-dom'

import './App.css'

import MainPage from './views/MainView';
import LoginPage from './views/LoginView';

class App extends Component {
  render() {
    return (
        <Router>
            <div>
                <ul>
                    <li><NavLink to="/">Home</NavLink></li>
                    <li><NavLink to="/login">Login</NavLink></li>
                </ul>
                <hl/>

                <Route exact path="/" component={MainPage}/>
                <Route path="/login" component={LoginPage}/>
            </div>
        </Router>
    );
  }
}

export default App;
