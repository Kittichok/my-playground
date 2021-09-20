import { expect } from 'chai';
import * as api from '../api/auth.js';

describe('Authentication', () => {
  it('should login fail', async () => {
    const body = {
      username: 'testssdnnsdxxxxx',
      password: 'admin',
    };
    const resp = await api.login(body);

    expect(resp.status).to.equal(400);
  });
});
