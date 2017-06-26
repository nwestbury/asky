/**
 * Created by James on 2017-05-13.
 */
import React, { Component } from 'react';
import { connect } from 'react-redux';
//import io from 'socket.io-client';
import { createSession } from '../actions/session';
import './LoginView.css';

import FacebookLogin from 'react-facebook-login';

//const socket = io("asky.ml");

const login = (facebookResp) => {
    //socket.emit("login", facebookResp, loginCallback);
};

const loginCallback = () => {} //socket.on("loginresp", responseServer);

const responseServer = (response) => {
    console.log(response);
};

const responseFacebook = (makeSession) => (response) => {
    console.log(response);
    makeSession(response.email, response.accessToken);
};

class LoginView extends Component {

    componentDidMount() {
        //loginCallback.connect();
    }

    componentWillUnmount() {
        //loginCallback.close();
    }

    render() {
        return (
            <div>
                <h2>Login Page</h2>
                <FacebookLogin
                    appId="315717428848841"
                    autoLoad={true}
                    fields="email"
                    callback={responseFacebook(this.props.makeSession)} />
            </div>
        );
    }
}

const mapDispatchToProps = dispatch => ({
    makeSession: (username, token) => dispatch(createSession(username, token)),
});

export default connect(null, mapDispatchToProps)(LoginView);