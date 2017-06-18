/**
 * Created by James on 2017-05-13.
 */

import React, { Component } from 'react';
import io from 'socket.io-client';
import './LoginView.css';

import FacebookLogin from 'react-facebook-login';

const socket = io("asky.ml");

const login = (facebookResp) => {
    socket.emit("login", facebookResp, loginCallback);
};

const loginCallback = socket.on("loginresp", responseServer);

const responseServer = (response) => {
    console.log(response);
};

const responseFacebook = (response) => {
    console.log(response);
};

class LoginView extends Component {

    componentDidMount() {
        loginCallback.connect();
    }

    componentDidUnmount() {
        loginCallback.close();
    }

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