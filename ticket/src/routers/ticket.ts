import { Router } from 'express';
import * as ticketService from '../services/ticketService';

const ticketRouter = Router();

ticketRouter.post('/create', ...ticketService.createValidatetor, ticketService.create);
ticketRouter.patch('/update', ...ticketService.updateValidatetor, ticketService.update);
ticketRouter.get('/list/:status*?', ticketService.getList);

export default ticketRouter;
