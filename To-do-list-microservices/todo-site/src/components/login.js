import React from "react"
import { navigate } from "gatsby"
import { handleLogin, isLoggedIn } from "../services/auth"

class Login extends React.Component {
  state = {
    username: ``,
    password: ``,
  }

  handleUpdate = event => {
    this.setState({
      [event.target.name]: event.target.value,
    })
  }

  handleSubmit = event => {
    event.preventDefault()
    handleLogin(this.state)
  }

  render() {
    if (isLoggedIn()) {
      navigate(`/app/todo`)
    }

    return (
      <form
        onSubmit={event => {
          this.handleSubmit(event)
          navigate(`/app/todo`)
        }}
      >
      <div className="p-4 max-w-xs mx-auto bg-white rounded-xl shadow-md m-3">
          <div className="mb-4">
            <label className="block text-grey-darker text-sm font-bold mb-2" htmlFor="username">
              Username
            </label>
            <input className="shadow appearance-none border rounded w-full py-2 px-3 text-grey-500" id="username" type="text" placeholder="Username" 
              name="username" 
              onChange={this.handleUpdate}/>
          </div>
          <div className="mb-6">
            <label className="block text-grey-darker text-sm font-bold mb-2" htmlFor="password">
              Password
            </label>
            <input className="shadow appearance-none border border-red rounded w-full py-2 px-3 text-grey-500 mb-3" id="password" type="password" placeholder="******************" 
              name="password"
              onChange={this.handleUpdate}/>
            {/* <p class="text-red-500 text-xs italic">{this.state.errormsg}</p> */}
          </div>
          <div className="flex items-center justify-between">
            <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" type="submit">
              Sign In
            </button>
          </div>
        </div>
      </form>
    )
  }
}

export default Login