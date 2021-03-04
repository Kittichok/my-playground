import React, { Component, FC } from 'react';
import { Route, useHistory } from 'react-router-dom';

function PrivateRoute({ component: Component, ...rest }: any) {

    let history = useHistory();
    return <Route {...rest} render={props => {
        // const currentUser = authenticationService.currentUserValue;
        // if (!currentUser) {
        //     history.push('login')
        // }

        return <Component {...props} />
    }} />
}

export default PrivateRoute