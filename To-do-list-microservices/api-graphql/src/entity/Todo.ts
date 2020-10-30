import {
  Entity,
  ObjectIdColumn,
  ObjectID,
  Column,
  PrimaryColumn,
  PrimaryGeneratedColumn,
  BaseEntity,
} from 'typeorm';
import { ObjectType, Field, ID } from 'type-graphql';

@Entity()
@ObjectType()
export class Todo extends BaseEntity {
  @Field(() => ID)
  @PrimaryGeneratedColumn()
  id: number;

  @Field(() => String)
  @Column()
  userID: string;

  @Field(() => String)
  @Column()
  text: string;

  @Field(() => Boolean)
  @Column({ default: false })
  isDone: boolean;
}
