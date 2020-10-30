import {Entity, ObjectIdColumn, ObjectID, Column, PrimaryColumn, PrimaryGeneratedColumn} from "typeorm";

@Entity()
export class Todo {

    @PrimaryColumn()
    UserID: number;

    @Column()
    Todo: string;

    @Column("boolean")
    isDone: boolean;
}
