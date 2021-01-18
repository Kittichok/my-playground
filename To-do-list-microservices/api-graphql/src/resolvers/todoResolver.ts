import { Resolver, Query, Mutation, Arg, Ctx } from 'type-graphql';
import { Todo } from '../entity/Todo';
import { CreateTodoInput } from '../inputs/createTodoInput';
import { UpdateTodoInput } from '../inputs/updateTodoInput';

@Resolver()
export class TodoResolver {
  @Query(() => [Todo])
  todoList(@Arg('userID', { nullable: true }) userID: string, @Ctx() ctx: any) {
    if (userID) return Todo.find({ where: { userID } });
    return Todo.find({ where: { userID: ctx.userid }});
  }

  @Query(() => Todo)
  todo(@Arg('userID') userID: string) {
    return Todo.findOne({ where: { userID } });
  }

  @Mutation(() => Todo)
  async createTodo(@Arg('text') text: string, @Ctx() ctx: any) {
    const todo = Todo.create({ text, userID: ctx.userid });
    await todo.save();
    return todo;
  }

  @Mutation(() => Todo)
  async updateTodo(@Arg('id') id: number, @Arg('data') data: UpdateTodoInput) {
    const todo = await Todo.findOne({ where: { id } });
    if (!todo) throw new Error('Todo not found!');
    Object.assign(todo, data);
    await todo.save();
    return todo;
  }

  @Mutation(() => Boolean)
  async deleteTodo(@Arg('id') id: string) {
    const todo = await Todo.findOne({ where: { id } });
    if (!todo) throw new Error('Todo not found!');
    await todo.remove();
    return true;
  }
}
