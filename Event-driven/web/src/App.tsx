import React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import './App.css';
import PrivateRoute from './components/privateRoute';
import NavBar from './components/navBar';
import Login from './pages/login';
import ProductList from './pages/productList'
import { route } from './config';

function App() {
  return (
    <div className="App">
      <Router>
        <Route exact path={route.login} component={Login} />
        <NavBar path={route.products} component={ProductList} title={"Products"}/>
        {/* <PrivateRoute path={route.myCart} component={MyCart} /> */}
      </Router>
    </div>
  );
}

export default App;
