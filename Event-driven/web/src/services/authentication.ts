import axios, { AxiosRequestConfig, AxiosResponse } from 'axios';

const login = async (email: string, password: string) => {
  const config: AxiosRequestConfig = {
    method: 'POST',
    url: `${process.env.REACT_APP_APIDOMAIN}/user/login`,
    data: { email, password },
  };

  return axios(config).then((resp: any) => {
    const { token } = resp.data;
    if (token) {
      localStorage.setItem('token', token);
      return token;
    }
    return;
  });
};

const logout = () => {
  localStorage.removeItem('token');
};

const register = async (email: string, password: string) => {
  const config: AxiosRequestConfig = {
    method: 'POST',
    url: `${process.env.REACT_APP_APIDOMAIN}/user/create`,
    data: { email, password },
  };

  return axios(config).then((resp: AxiosResponse) => {
    return resp.status === 201;
  });
};

export const authenticationService = {
  login,
  logout,
  register,
  token: localStorage.getItem('token') as string,
};
