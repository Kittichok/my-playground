import React from "react";
import {
  BrowserRouter as Router,
  Route
} from "react-router-dom";
import './App.css'
import PrivateRoute from "./components/privateRoute";
import CreateParty from "./pages/createParty";
import Login from './pages/login';
import PartyList from "./pages/partyList";
import Register from "./pages/register";

function App() {
  return (
    <div className="App">
      <Router>
        <Route exact path="/" component={Login} />
        <Route path="/register" component={Register} />
        <PrivateRoute path="/partyList" component={PartyList} />
        <PrivateRoute path="/createParty" component={CreateParty} />
      </Router>
    </div>
  );
}

export default App;
