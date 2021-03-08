import { Router } from 'express';
import { verifyJWT } from './middleware';
import partyRouter from './party';
import userRouter from './users';

const routes = Router();

routes.use('/user', userRouter);
routes.use('/party', verifyJWT, partyRouter);

export default routes;
