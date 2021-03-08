import { Request, Response, Router } from 'express';
import { User } from '../entity/User';
import { getManager } from 'typeorm';
import jwt from 'jsonwebtoken';
import * as utils from '../utils';

const userRouter = Router();

//TODO body validation
userRouter.post('/login', async (req: Request, res: Response) => {
  const { email, password } = req.body;

  const userRepository = getManager().getRepository(User);

  const existUser = await userRepository.findOne({ email });
  if (existUser) {
    const correctPass = utils.validateHashPassword(password, existUser.saltKey, existUser.passward);
    if (correctPass) {
      const token = generateToken(existUser);
      return res.json({
        token,
      });
    }
  }
  return res.status(400).send();
});

//TODO body validation
userRouter.post('/create', async (req: Request, res: Response) => {
  const { email, password } = req.body;

  const userRepository = getManager().getRepository(User);

  const { hash, salt } = utils.saltHashPassword(password);
  const r = await userRepository.insert({
    email: email,
    passward: hash,
    saltKey: salt,
  });

  if (r.raw != 0) {
    return res.status(201).send();
  }
  return res.status(400).send();
});

const generateToken = (user: User) => {
  const oneDay = Math.floor(Date.now() / 1000) + 3600 * 60;
  const token = jwt.sign(
    {
      exp: oneDay,
      data: user.email,
    },
    process.env.JWTKEY,
  );
  return token;
};

export default userRouter;
