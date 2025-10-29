import { Router } from 'express';
import studentsRouter from './students';
import collegesRouter from './colleges';

const router = Router();

router.use('/students', studentsRouter);
router.use('/colleges', collegesRouter);

export default router;
