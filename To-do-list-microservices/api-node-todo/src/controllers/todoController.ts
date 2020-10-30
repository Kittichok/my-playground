import { Request, Response } from 'express';
import { Todo } from '../entity/Todo';
import { getManager } from 'typeorm';

export const create = async (req: Request, res: Response) => {
  const text = req.body.todo;
  
  const id = req["userid"];
  const todoRepository = getManager().getRepository(Todo);

  const r = await todoRepository.insert({
    isDone: false,
    Todo: text,
    UserID: id,
  });

  if (r.raw != 0) {
    res.status(201).send();
  } 
  res.status(400).send();
};
