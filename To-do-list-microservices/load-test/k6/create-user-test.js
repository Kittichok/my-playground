import http from 'k6/http';
import { check, group, sleep } from 'k6';


export let options = {
  stages: [
    { duration: '1m', target: 10 },
    { duration: '1m', target: 20 },
    { duration: '1m', target: 0 }, 
    // { duration: '5m', target: 100 }, // simulate ramp-up of traffic from 1 to 100 users over 5 minutes.
    // { duration: '10m', target: 100 }, // stay at 100 users for 10 minutes
    // { duration: '5m', target: 0 }, // ramp-down to 0 users
  ],
  thresholds: {
    http_req_duration: ['p(99)<1500'], // 99% of requests must complete below 1.5s
    'logged in successfully': ['p(99)<1500'], // 99% of requests must complete below 1.5s
  },
  // noConnectionReuse: true,
};

const data = JSON.parse(open('./data.json'));
const BASE_URL = 'http://auth.com';
let index = 0;

export default () => {
  let user = data.users[index];
  if(!user) { return; }
  console.log(user.username)
  console.log(user.password)
  const payload = JSON.stringify({
    username: user.username,
    password: user.password
  });
  // const params = {
  //   headers: {
  //     'Content-Type': 'application/json',
  //   },
  // };
  // let loginRes = http.post(`${BASE_URL}/v1/api/signin`, payload, params);
  // let accessToken = "";

  // check(loginRes, {
  //   'logged in successfully': (resp) => resp.json('AccessToken') !== '',
  // });


  const signupRes = http.post(`${BASE_URL}/v1/api/signup`, payload);

  check(signupRes, {
    'signup in successfully': (resp) => resp.status == 200,
  });

  // const accessToken = signupRes.json('AccessToken')
  // if (loginRes.status != 200) {
  // }

  index++;
  sleep(1);
};
