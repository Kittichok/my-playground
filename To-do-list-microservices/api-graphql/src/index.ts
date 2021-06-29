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
import cors from 'cors'
import { initTracer } from "./utils/traceLogging";

const tracer = initTracer("todo-service");

const PORT = process.env.PORT || 4000

async function main() {
  const connection = await createConnection();
  await connection.synchronize();
  const schema = await buildSchema({
    resolvers: [TodoResolver],
  });
  const app = express();
  const server = new ApolloServer({
    schema,
    context: ({ req }) => {
      const context = {
        req,
        userid: req.body.userid,
      };
      return context;
    },
  });
  const path = '/graphql';
  app.use(cors())
  app.use(bodyParser.json())
  app.get('/ping', (_req:any, res:any) => {
    const span = tracer.startSpan("GET /ping");
    res.status(200).send('pong')
    span.finish();
  });
  app.use(path, verifyJWT);
  server.applyMiddleware({ app, path });
  app.listen(PORT, () => console.log('Server has started!', PORT));
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
      // console.log('decoded',decoded);
      req.body["userid"] = decoded.jti;
      next();
    } else {
      const span = tracer.startSpan("token unauthorize");
      span.finish();
      res.status(401).send();
    }
  }
};