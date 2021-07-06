import connection from '../services/dbService';
import * as request from 'supertest';
import server from '../server';
import * as assert from 'assert';

beforeAll(async () => {
	await connection.create();
});

afterAll(async () => {
	await connection.close();
});

beforeEach(async () => {
	await connection.clear();
});

it('should get error missing body', async () => {
	request(server)
		.post('/ticket/create')
		.send({
			description: 'test',
		})
		.expect(400);
});

it('should get 201', async () => {
	request(server)
		.post('/ticket/create')
		.send({
			title: 'Task1',
			description: 'test',
		})
		.expect(201);
});

it('should update success', async () => {
	await request(server)
		.post('/ticket/create')
		.send({
			title: 'Task1',
			description: 'test',
			contract: 'test',
		})
		.expect(201);

	request(server)
		.patch('/ticket/update')
		.send({
			description: 'change something',
			status: 'accepted',
		})
		.expect(200);
});

it('should update fail', async () => {
	await request(server)
		.post('/ticket/create')
		.send({
			title: 'Task1',
			description: 'test',
			contract: 'test',
		})
		.expect(201);

	request(server)
		.patch('/ticket/update')
		.send({
			status: 'wrong wrong',
		})
		.expect(400);
});

it('should have list', async () => {
	await request(server)
		.post('/ticket/create')
		.send({
			title: 'Task2',
			description: 'test',
			contract: 'test',
		})
		.expect(201);

	request(server).get('/ticket/list').then((response) => {
		assert(response.body[0].title, 'Task2');
	});
});
