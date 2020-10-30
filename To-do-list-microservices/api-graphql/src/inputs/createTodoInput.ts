import { InputType, Field } from "type-graphql";

@InputType()
export class CreateTodoInput {
  @Field()
  userID: string;
  @Field()
  text: string;
}