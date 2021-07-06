import * as express from 'express';
import router from './routers';
import * as bodyParser from 'body-parser';

const server = express();

server.use(bodyParser.urlencoded({ extended: false }));
server.use(bodyParser.json());

server.use('/ping', (req, res) => {
	res.send('pong');
});

server.use(router);

export default server;
