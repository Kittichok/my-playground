import { expect } from 'chai';
import Faker from 'faker';
import * as api from '../api/index.js';

describe('Register', () => {
  const username = 'user_' + Faker.datatype.number();
  it('should register successful', async () => {
    const body = {
      username,
      password: 'admin',
    };
    const resp = await api.register.register(body);

    expect(resp.status).to.equal(201);
  });
  it('should login successful', async () => {
    const body = {
      username,
      password: 'admin',
    };
    const resp = await api.auth.login(body);

    expect(resp.status).to.equal(200);
  });
});
