import { Entity, PrimaryGeneratedColumn, Column } from 'typeorm';

@Entity()
export class Ticket {
	@PrimaryGeneratedColumn() id: number;
	@Column() title: string;
	@Column() description: string;
	@Column() contract: string;
	@Column() createAt: Date;
	@Column() updateAt: Date;
	@Column() status: ticketStatus;
}

export enum ticketStatus {
	pending = 'pending',
	accepted = 'accepted',
	rejected = 'rejected',
	resolved = 'resolved',
}
