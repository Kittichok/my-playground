import React from 'react';
import { Switch, Route, BrowserRouter } from 'react-router-dom';
import { Login } from './Container/Login';
import { Todo } from './Container/Todo';
import styled from 'styled-components';
import 'antd/dist/antd.css';

const Center = styled.div`
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
`;

function App() {
  return (
    <Center>
      <BrowserRouter>
        <Switch>
          <Route exact path="/" component={Login} />
          <Route exact path="/todo" component={Todo} />
        </Switch>
      </BrowserRouter>
    </Center>
  );
}

export default App;
