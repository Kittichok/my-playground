import { Resolver, Query, Mutation, Arg, Ctx } from 'type-graphql';
import { Todo } from '../entity/Todo';
import { CreateTodoInput } from '../inputs/createTodoInput';
import { UpdateTodoInput } from '../inputs/updateTodoInput';
import { initTracer } from "../utils/traceLogging";

const tracer = initTracer("todo-service");

@Resolver()
export class TodoResolver {
  @Query(() => [Todo])
  todoList(@Arg('userID', { nullable: true }) userID: string, @Ctx() ctx: any) {
    if (userID) return Todo.find({ where: { userID } });
    const span = tracer.startSpan("get todoList");
    span.log({ res: "userID: " + userID});
    span.finish();
    return Todo.find({ where: { userID: ctx.userid }});
  }

  @Query(() => Todo)
  todo(@Arg('userID') userID: string) {
    const span = tracer.startSpan("get todo");
    span.log({ res: "userID: " + userID});
    span.finish();
    return Todo.findOne({ where: { userID } });
  }

  @Mutation(() => Todo)
  async createTodo(@Arg('text') text: string, @Ctx() ctx: any) {
    const todo = Todo.create({ text, userID: ctx.userid });
    await todo.save();
    const span = tracer.startSpan("create todo");
    span.log({ res: "userID: " + ctx.userid});
    span.finish();
    return todo;
  }

  @Mutation(() => Todo)
  async updateTodo(@Arg('id') id: number, @Arg('data') data: UpdateTodoInput) {
    const todo = await Todo.findOne({ where: { id } });
    if (!todo) throw new Error('Todo not found!');
    Object.assign(todo, data);
    await todo.save();
    const span = tracer.startSpan("update todo");
    span.log({ res: "todo id: " + id});
    span.finish();
    return todo;
  }

  @Mutation(() => Boolean)
  async deleteTodo(@Arg('id') id: string) {
    const todo = await Todo.findOne({ where: { id } });
    if (!todo) throw new Error('Todo not found!');
    await todo.remove();
    const span = tracer.startSpan("delete todo");
    span.log({ res: "todo id: " + id});
    span.finish();
    return true;
  }
}
