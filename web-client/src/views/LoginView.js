/**
 * Created by James on 2017-05-13.
 */

import React, { Component } from 'react';
import './LoginView.css';

import FacebookLogin from 'react-facebook-login';

const responseFacebook = (response) => {
    console.log(response);
};

class LoginView extends Component {
    render() {
        return (
            <div>
                <h2>Login Page</h2>
                <FacebookLogin
                    appId="315717428848841"
                    autoLoad={true}
                    fields="email"
                    callback={responseFacebook} />
            </div>
        );
    }
}

export default LoginView;