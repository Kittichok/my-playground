import 'reflect-metadata';
import { createConnection } from 'typeorm';
import { buildSchema } from 'type-graphql';
import { ApolloServer } from 'apollo-server';
import { TodoResolver } from './resolvers/todoResolver';

//TODO wrap with express for use a jwt authorize
async function main() {
  const connection = await createConnection();
  await connection.synchronize();
  const schema = await buildSchema({
    resolvers: [TodoResolver],
  });
  const server = new ApolloServer({ schema });
  // server.applyMiddleware()
  await server.listen(4000);
  console.log('Server has started!');
}

main();
