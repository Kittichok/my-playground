import React from 'react';
import { Route, useHistory } from 'react-router-dom';
import { authenticationService } from '../services/authentication';

function PrivateRoute({ component: Component, ...rest }: any) {
  let history = useHistory();
  return (
    <Route
      {...rest}
      render={(props) => {
        const token = authenticationService.token;
        if (!token) {
          history.push('/');
        }

        return <Component {...props} />;
      }}
    />
  );
}

export default PrivateRoute;
