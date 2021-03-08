import { Request, Response, Router } from 'express';
import { getManager } from 'typeorm';
import { Party } from '../entity/Party';

const partyRouter = Router();

//TODO Limit max
partyRouter.get('/list', async (_req: Request, res: Response) => {
  const partyRepository = getManager().getRepository(Party);

  const partyList = await partyRepository.find();

  if (partyList) {
    return res.json(partyList);
  }
  return res.status(400).send();
});

//TODO add owner id
partyRouter.post('/create', async (req: Request, res: Response) => {
  const { name, totalMember } = req.body;

  const partyRepository = getManager().getRepository(Party);

  const r = await partyRepository.insert({
    imgUrl: '',
    name,
    totalMember,
    currentMember: 0,
  });

  if (r.raw != 0) {
    return res.status(201).send();
  }
  return res.status(400).send();
});

//TODO list member in party
partyRouter.get('/join/:partyID', async (req: Request, res: Response) => {
  const partyID = +req.params.partyID;
  const partyRepository = getManager().getRepository(Party);

  const party = await partyRepository.findOne({ id: partyID });
  party.currentMember++;
  await partyRepository.save(party);
  return res.status(201).send();
});

export default partyRouter;
