import { randomIntBetween, randomString } from 'https://jslib.k6.io/k6-utils/1.1.0/index.js';
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
const endpoint = `${BASE_URL}/product/api/v1/product`
// let index = 0;

export default () => {
  // let product = data.products[index];
  // if(!product) { return; }

  const randomName = randomString(8);
  const randomPrice = randomIntBetween(50, 5000);
  const randomQuantity = randomIntBetween(100, 5000);
  const payload = JSON.stringify({
    name: randomName,
    quantity: randomPrice,
    price: randomQuantity
  });
  const resp = http.post(endpoint, payload);

  check(resp, {
    'create product successfully': (resp) => resp.status == 201,
  });

  // index++;
  sleep(1);
};
