import axios from 'axios';
import { config } from './config.js';

const register = async body => {
  const url = `${config.url}/v1/signup`;
  const resp = await axios.post(url, body).then(response => response);
  return resp;
};

export { register };
