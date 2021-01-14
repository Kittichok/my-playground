import * as React from "react"
import { Router } from "@reach/router"
import Layout from "../components/layout"
import Login from "../components/login"
import Todo from "../components/todo"
import PrivateRoute from "../components/privateRoute"

// styles


// markup
// TODO : add route login page and app page
const App = () => {
  return (
    <Layout>
      <Router>
        <PrivateRoute path="/app/todo" component={Todo} />
        <Login path="/app/login" />
      </Router>
    </Layout>
  )
}

export default App
