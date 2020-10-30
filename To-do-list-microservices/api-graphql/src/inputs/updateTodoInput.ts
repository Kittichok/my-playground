import { InputType, Field } from "type-graphql";

@InputType()
export class UpdateTodoInput {
  @Field({ nullable: true })
  text: string;
  @Field({ nullable: true })
  isDone?: boolean;
}