import axios from 'axios';
import { config } from './config.js';

const login = async body => {
  const url = `${config.url}/v1/signin`;
  const resp = await axios
    .post(url, body)
    .then(response => response)
    .catch(err => err.response);
  return resp;
};

export { login };
