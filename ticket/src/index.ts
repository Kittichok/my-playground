import 'reflect-metadata';
import { createConnection } from 'typeorm';
import server from './server';

const port = process.env.PORT || 3000;

createConnection()
	.then(async (_connection) => {
		server.listen(port, () => {
			console.log('listening on port ' + port);
		});
	})
	.catch((error) => console.log(error));
