import 'reflect-metadata';
import { createConnection } from 'typeorm';
import server from './server';

createConnection()
  .then(async (_connection) => {
    const port = process.env.PORT;
    server.listen(port, () => {
      console.log('  App is running at http://localhost:%d in %s mode', port, server.get('env'));
      console.log('Press CTRL-C to stop\n');
    });
  })
  .catch((error) => console.log(error));
