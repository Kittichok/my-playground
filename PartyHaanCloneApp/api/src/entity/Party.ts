import {
  Entity,
  ObjectIdColumn,
  ObjectID,
  Column,
  PrimaryColumn,
  PrimaryGeneratedColumn,
} from 'typeorm';

@Entity()
export class Party {
  @PrimaryGeneratedColumn('increment')
  id: number;

  @Column()
  name: string;

  @Column()
  currentMember: number;

  @Column()
  totalMember: number;

  @Column()
  imgUrl: string;
}
