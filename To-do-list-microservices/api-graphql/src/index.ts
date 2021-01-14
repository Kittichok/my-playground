import 'reflect-metadata';
import { createConnection } from 'typeorm';
import { buildSchema } from 'type-graphql';
// import { ApolloServer } from 'apollo-server';
import { TodoResolver } from './resolvers/todoResolver';
import express from 'express';
import { Request, Response, NextFunction } from 'express';
import bodyParser from 'body-parser';
import { ApolloServer } from 'apollo-server-express';
import * as jwt from 'jsonwebtoken';

const PORT = process.env.PORT || 4000

//TODO wrap with express for use a jwt authorize
async function main() {
  const connection = await createConnection();
  await connection.synchronize();
  const schema = await buildSchema({
    resolvers: [TodoResolver],
  });
  const app = express();
  const server = new ApolloServer({
    schema
  });
  const path = '/graphql';
  app.use(bodyParser.json())
  app.use(path, verifyJWT);
  server.applyMiddleware({ app, path });
  app.listen(PORT, () => console.log('Server has started!'));
}

main();


const verifyJWT = (req: Request, res: Response, next: NextFunction) => {
  if (
    req.headers.authorization &&
    req.headers.authorization.split(' ')[0] === 'Bearer'
  ) {
    const token = req.headers.authorization.split(' ')[1];
    console.log(token);
    
    const decoded = jwt.verify(token, 'AllYourBase') as any;
    if (decoded) {
      req.body["userid"] = decoded.jti;
      next();
    } else {
      res.status(401).send();
    }
  }
};