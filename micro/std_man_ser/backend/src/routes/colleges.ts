import { Router } from "express";
import prisma from "../prisma";

const router = Router();

router.get("/", async (req, res, next) => {
  try {
    const colleges = await prisma.college.findMany({
      include: { departments: true, programs: true },
    });
    res.json({ ok: true, data: colleges });
  } catch (err) {
    next(err);
  }
});

export default router;
