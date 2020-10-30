import 'reflect-metadata';
import { createConnection } from 'typeorm';
import { Todo } from './entity/Todo';
import server from './server';

createConnection()
  .then(async (connection) => {
    // console.log('Inserting a new user into the database...');
    // const user = new User();
    // user.firstName = 'Timber';
    // user.lastName = 'Saw';
    // user.age = 25;
    // await connection.manager.save(user);
    // console.log('Saved a new user with id: ' + user.id);

    console.log('Loading todos from the database...');
    const todos = await connection.manager.find(Todo);
    console.log('Loaded todos: ', todos);

    const port = 3000;
    server.listen(port, () => {
      console.log(
        '  App is running at http://localhost:%d in %s mode',
        port,
        server.get('env')
      );
      console.log('  Press CTRL-C to stop\n');
    });
  })
  .catch((error) => console.log(error));
