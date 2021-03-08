import React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import './App.css';
import PrivateRoute from './components/privateRoute';
import CreateParty from './pages/createParty';
import Login from './pages/login';
import PartyList from './pages/partyList';
import Register from './pages/register';
import { route } from './config';

function App() {
  return (
    <div className="App">
      <Router>
        <Route exact path={route.login} component={Login} />
        <Route path={route.register} component={Register} />
        <PrivateRoute path={route.listing} component={PartyList} />
        <PrivateRoute path={route.createParty} component={CreateParty} />
      </Router>
    </div>
  );
}

export default App;
