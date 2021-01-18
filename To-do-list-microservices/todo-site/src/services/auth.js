import { login } from './api'

export const isBrowser = () => typeof window !== "undefined"

export const getUser = () =>
  isBrowser() && window.localStorage.getItem("gatsbyUser")
    ? JSON.parse(window.localStorage.getItem("gatsbyUser"))
    : {}

const setUser = user =>
  window.localStorage.setItem("gatsbyUser", JSON.stringify(user))

export const handleLogin = async ({ username, password }) => {
  const resp = await login({ username, password })
  if(resp) {
    return setUser({
      token: resp.AccessToken,
    })
  }
  return false
}

export const isLoggedIn = () => {
  const user = getUser()

  return user && !!user.token
}

export const logout = callback => {
  setUser({})
  callback()
}