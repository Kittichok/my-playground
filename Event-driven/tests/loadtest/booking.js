import { randomIntBetween } from 'https://jslib.k6.io/k6-utils/1.1.0/index.js';
import { check, sleep } from 'k6';
import http from 'k6/http';

export let options = {
  stages: [
    { duration: '2m', target: 30 },
    { duration: '2m', target: 30 },
    { duration: '1m', target: 0 },
    // { duration: '5m', target: 100 }, // simulate ramp-up of traffic from 1 to 100 users over 5 minutes.
    // { duration: '10m', target: 100 }, // stay at 100 users for 10 minutes
    // { duration: '5m', target: 0 }, // ramp-down to 0 users
  ],
  thresholds: {
    http_req_duration: ['p(99)<1500'], // 99% of requests must complete below 1.5s
    // 'logged in successfully': ['p(99)<1500'], // 99% of requests must complete below 1.5s
  },
  noConnectionReuse: true,
};

// const data = JSON.parse(open('./data.json'));
const BASE_URL = 'http://localhost:8000';
const endpoint = `${BASE_URL}/booking/api/v1/booking`
const endpointSubmit = (id) => `${endpoint}/${id}/submit`
// let index = 0;

export default () => {
  // let product = data.products[index];
  // if(!product) { return; }

  const randomID = randomIntBetween(1, 5000);
  const randomProductID = randomIntBetween(1, 9999);
  const randomQuantity = randomIntBetween(1, 50);
  const randomPrice = randomIntBetween(1, 400);
  const payload = JSON.stringify({
    user_id: randomID,
    products: [
        {
          product_id: randomProductID,
          quantity: randomQuantity,
          price: randomPrice
        }
    ]
  });
  const resp = http.post(endpoint, payload);

  // console.log(JSON.stringify(resp.body))
  check(resp, {
    'create booking successfully': (resp) => resp.status == 201,
  });

  const respSubmit = http.get(endpointSubmit(JSON.parse(resp.body).ID), payload);

  check(respSubmit, {
    'submit booking successfully': (respSubmit) => respSubmit.status == 202,
  });

  sleep(1);
};
