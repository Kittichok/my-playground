import { Request, Response } from 'express';
import { body, check, CustomValidator, validationResult } from 'express-validator';
import { createQueryBuilder, FindConditions, getManager } from 'typeorm';
import { Ticket, ticketStatus } from '../entity/ticket';

const create = async (req: Request, res: Response) => {
	try {
		const errors = validationResult(req);

		if (!errors.isEmpty()) {
			res.status(400).json({ errors: errors.array() });
			return;
		}

		const { title, description, contract } = req.body;

		const ticketRepository = getManager().getRepository(Ticket);
		const ticket = {
			title,
			description,
			contract,
			createAt: new Date(),
			updateAt: new Date(),
			status: 'pending',
		} as Ticket;

		const createdTicket = await ticketRepository.insert(ticket);

		if (createdTicket) {
			res.sendStatus(201);
			return;
		}
		res.sendStatus(500);
		return;
	} catch (err) {
		console.error(err);
		res.sendStatus(500);
		return;
	}
};

const update = async (req: Request, res: Response) => {
	try {
		const errors = validationResult(req);

		if (!errors.isEmpty()) {
			res.status(400).json({ errors: errors.array() });
			return;
		}

		const { title, description, contract, status, id } = req.body;

		const ticketRepository = getManager().getRepository(Ticket);
		const existingTicket = await ticketRepository.findOne({ id });

		if (!existingTicket) {
			res.sendStatus(400);
		}

		existingTicket.title = title || existingTicket.title;
		existingTicket.description = description || existingTicket.description;
		existingTicket.contract = contract || existingTicket.contract;
		existingTicket.status = status || existingTicket.status;
		existingTicket.updateAt = new Date();

		await ticketRepository.update({ id }, existingTicket);

		res.sendStatus(200);
	} catch (err) {
		console.error(err);
		res.sendStatus(500);
	}
};

const getList = async (req: Request, res: Response) => {
	const { status } = req.params;

	const condition: FindConditions<Ticket> = {};

	if (status) {
		condition.status = status as ticketStatus;
	}
	const tickets = await createQueryBuilder(Ticket).where(condition).orderBy('ticket.updateAt', 'DESC').getMany();

	res.json(tickets);
};

const validStatus: CustomValidator = (value: ticketStatus) => {
	return Object.values(ticketStatus).includes(value);
};
const createValidatetor = [ check('title').exists(), check('contract').exists(), check('description').exists() ];
const updateValidatetor = [ check('status').custom(validStatus) ];

export { create, update, getList, createValidatetor, updateValidatetor };
