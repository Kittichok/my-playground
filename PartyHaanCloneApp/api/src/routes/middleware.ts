import { Request, Response, NextFunction } from 'express';
import * as jwt from 'jsonwebtoken';

export const verifyJWT = (req: Request, res: Response, next: NextFunction) => {
  if (req.headers.authorization && req.headers.authorization.split(' ')[0] === 'Bearer') {
    const token = req.headers.authorization.split(' ')[1];
    // console.log(token);

    const decoded = jwt.verify(token, process.env.JWTKEY) as any;
    if (decoded) {
      next();
    } else {
      res.status(401).send();
    }
  }
};
