import axios, { AxiosRequestConfig, AxiosResponse } from 'axios';
import { authenticationService } from './authentication';

const getList = async () => {
  const config: AxiosRequestConfig = {
    method: 'GET',
    url: `${process.env.REACT_APP_APIDOMAIN}/party/list`,
    headers: {
      Authorization: 'Bearer ' + authenticationService.token,
    },
  };

  return axios(config).then((resp: AxiosResponse<[]>) => {
    const partyList = resp.data;
    return partyList;
  });
};

const create = async (partyName: string, totalMember: number) => {
  const config: AxiosRequestConfig = {
    method: 'POST',
    url: `${process.env.REACT_APP_APIDOMAIN}/party/create`,
    headers: {
      Authorization: 'Bearer ' + authenticationService.token,
    },
    data: {
      name: partyName,
      totalMember,
    },
  };

  return axios(config).then((resp: AxiosResponse) => {
    return resp.status === 201;
  });
};

const join = async (partyID: number) => {
  const config: AxiosRequestConfig = {
    method: 'GET',
    url: `${process.env.REACT_APP_APIDOMAIN}/party/join/${partyID}`,
    headers: {
      Authorization: 'Bearer ' + authenticationService.token,
    },
  };

  return axios(config).then((resp: AxiosResponse) => {
    return resp.status === 201;
  });
};

export { getList, create, join };
