import axios, { AxiosRequestConfig, AxiosResponse } from 'axios';
import { authenticationService } from './authentication';

const getList = async () => {
  const config: AxiosRequestConfig = {
    method: 'GET',
    url: `${process.env.REACT_APP_APIDOMAIN}/api/v1/products`,
    headers: {
      Authorization: 'Bearer ' + authenticationService.token,
    },
  };

  return axios(config).then((resp: AxiosResponse<any>) => {
    const productList = resp.data.product;
    return productList;
  });
};

export { getList };
