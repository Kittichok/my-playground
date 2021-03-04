import {
  BrowserRouter as Router,
  Route
} from "react-router-dom";
import './App.css'
import Login from './pages/login';
import Register from "./pages/register";

function App() {
  return (
    <div className="App">
      <Router>
        <Route exact path="/" component={Login}/>
        <Route path="/register" component={Register}/>
        {/* <Authen path="/partyList" component={}/> */}
        {/* <Authen path="/createParty" component={}/> */}
      </Router>
    </div>
  );
}

export default App;
