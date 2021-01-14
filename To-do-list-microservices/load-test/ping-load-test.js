import http from 'k6/http';
import { check, group, sleep } from 'k6';

export let options = {
  stages: [
    { duration: '1m', target: 20 },
    { duration: '1m', target: 0 },
  ],
  thresholds: {
    requests: ['count < 100'],
  },
  // noConnectionReuse: true,
};

const BASE_URL = 'http://localhost:5000';
const USERNAME = 'test';
const PASSWORD = 'test';
const API_URL = 'http://localhost:4000';

export default () => {
  const res = http.get(`${BASE_URL}/ping`);
  sleep(1);
  const checkRes = check(res, {
    'status is 200': (r) => r.status === 200,
    'response body': (r) => r.body.indexOf('pong') !== -1,
  });
};
