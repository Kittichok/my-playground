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
const BASE_URL = 'http://localhost:5000';
// const USERNAME = 'test';
// const PASSWORD = 'test';
const API_URL = 'http://localhost:4000';
let index = 0;

export default () => {
  let user = data.users[index];
  let loginRes = http.post(`${BASE_URL}/v1/api/signin/`, {
    username: user.username,
    password: user.password,
  });

  check(loginRes, {
    'logged in successfully': (resp) => resp.json('AccessToken') !== '',
  });

  let authHeaders = {
    headers: {
      Authorization: `Bearer ${loginRes.json('AccessToken')}`,
    },
  };
  let query = `
  query {
    todoList(userID: "1150") {
      text,
      isDone
    }
  }`;

  let myObjects = http.post(`${API_URL}/graphql/`, JSON.stringify({ query: query }), authHeaders).json('data');
  // console.log(myObjects)
  check(myObjects, { 'retrieved todo list': (obj) => obj && obj.todoList && obj.todoList.length > 0 });
  index++;
  sleep(1);
};
