require('dotenv').config();
import * as express from 'express';
import * as bodyParser from 'body-parser';
import * as cors from 'cors';
import * as todoController from './controllers/todoController';
import { verifyJWT } from './controllers/middleware';

const app = express();
app.use(bodyParser());
app.use(cors());

// respond with "hello world" when a GET request is made to the homepage
app.get('/', function (_req, res) {
  res.send('hello world');
});
// TODO middleware jwt
// TODO api todo create, getList, update, delete
app.get('/create', verifyJWT, todoController.create);

export default app;
