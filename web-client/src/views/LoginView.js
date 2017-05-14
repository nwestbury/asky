/**
 * Created by James on 2017-05-13.
 */

import React, { Component } from 'react';
import './LoginView.css';

class LoginView extends Component {
    render() {
        return (
            <div>
                <h2>Login Page</h2>
                <div className="field">
                    <span className="label">Username: </span><input name="username"/>
                </div>
                <br/>
                <div className="field">
                    <span className="label">Password: </span><input name="password"/>
                </div>
            </div>
        );
    }
}

export default LoginView;